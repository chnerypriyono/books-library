package main

import (
	"net/http"
	"database/sql"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"
    _ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/mux"
)

const (
	dbDriver = "mysql"
)

type BookOverview struct {
    id    	int
    title  	string
    author 	string   
  }

type BookDetail struct {
	id    		int
	title  		string
	author 		string
	description string
	rating 		int
}

func (app *application) getBooksHandler(w http.ResponseWriter, r *http.Request) {
    db, err := sql.Open(dbDriver, os.Getenv(MYSQL_PRIVATE_URL))
    if err != nil {
      panic(err.Error())
    }
    defer db.Close()    

    books, err := GetBooks(db)
    if err != nil {
      http.Error(w, "Not found", http.StatusNotFound)
      return
    }
   
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}

func (app *application) getBookDetailHandler(w http.ResponseWriter, r *http.Request) {
    db, err := sql.Open(dbDriver, os.Getenv(MYSQL_PRIVATE_URL))
    if err != nil {
      panic(err.Error())
    }
    defer db.Close()
    
    vars := mux.Vars(r)
    idStr := vars["id"]

    bookID, err := strconv.Atoi(idStr)

    book, err := GetBookDetail(db, bookID)
    if err != nil {
      http.Error(w, "Book not found", http.StatusNotFound)
      return
    }
   
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(book)
}


func GetBooks(db *sql.DB) ([]BookOverview, error) {
    query := "SELECT id, title, author FROM books"
    rows, err := db.Query(query)

    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var books []BookOverview

    for rows.Next() {
        var book BookOverview
        if err := rows.Scan(&book.id, &book.title, &book.author); err != nil {
            return books, err
        }
        books = append(books, book)
    }
    if err = rows.Err(); err != nil {
        return books, err
    }
    return books, nil
}


func GetBookDetail(db *sql.DB, id int) (*BookDetail, error) {
    query := "SELECT * FROM books WHERE id = ?"
    row := db.QueryRow(query, id)

    book := &BookDetail{}
    err := row.Scan(&book.id, &book.title, &book.author, &book.description, &book.rating)
    if err != nil {
        return nil, err
    }
    return book, nil
}