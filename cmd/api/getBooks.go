package main

import (
	"net/http"
	"database/sql"
    "encoding/json"       
    "strconv"
    _ "github.com/lib/pq"   
    "os"
)

const (
	dbDriver = "postgres"
)

type BookOverview struct {
    Id    	int 	`json:"id"`
    Title  	string	`json:"title"`
    Author 	string	`json:"author"`
  }

type BookDetail struct {
	Id    		int 	`json:"id"`
    Title  		string	`json:"title"`
    Author 		string	`json:"author"`
	Description string	`json:"description"`
	Rating 		int 	`json:"rating"`
}

func (app *application) getBooksHandler(w http.ResponseWriter, r *http.Request) {
    db, err := sql.Open(dbDriver, os.Getenv("DATABASE_URL"))
    if err != nil {
      panic(err.Error())
    }
    defer db.Close()    

    books, err := GetBooks(app, db)
    if err != nil {
      app.logger.Error(err.Error())
      http.Error(w, "Not found", http.StatusNotFound)
      return
    }
   
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}

func (app *application) getBookDetailHandler(w http.ResponseWriter, r *http.Request) {
    db, err := sql.Open(dbDriver, os.Getenv("DATABASE_URL"))
    if err != nil {
      panic(err.Error())
    }
    defer db.Close()
    
    idStr := r.URL.Query().Get("id")

    bookID, err := strconv.Atoi(idStr)

    book, err := GetBookDetail(app, db, bookID)
    if err != nil {
      app.logger.Error(err.Error())
      http.Error(w, "Book not found", http.StatusNotFound)
      return
    }
   
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(book)
}


func GetBooks(app *application, db *sql.DB) ([]BookOverview, error) {
    query := "SELECT id, title, author FROM books;"
    rows, err := db.Query(query)

    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var books []BookOverview

    for rows.Next() {
        var book BookOverview
        if err := rows.Scan(&book.Id, &book.Title, &book.Author); err != nil {
            return books, err
        }
        app.logger.Info("book: ", book)
        books = append(books, book)
    }
    if err = rows.Err(); err != nil {    	
        return books, err
    }
    app.logger.Info("books: ", books)
    return books, nil
}


func GetBookDetail(app *application, db *sql.DB, id int) (*BookDetail, error) {
	app.logger.Info("id: ", id)
    query := "SELECT * FROM books WHERE id = $1;"
    row := db.QueryRow(query, id)

    book := &BookDetail{}
    err := row.Scan(&book.Id, &book.Title, &book.Author, &book.Description, &book.Rating)
    if err != nil {    	
        return nil, err
    }
    return book, nil
}