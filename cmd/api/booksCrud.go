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

    books, err := getBooks(app, db)
    if err != nil {
      app.logger.Error(err.Error())
      http.Error(w, "Not found", http.StatusNotFound)
      return
    }
   
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}

func getBooks(app *application, db *sql.DB) ([]BookOverview, error) {
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
        app.logger.Info("retrieved book row", "book", book)
        books = append(books, book)
    }
    if err = rows.Err(); err != nil {    	
        return books, err
    }
    app.logger.Info("retrieved books", "books", books)
    return books, nil
}

func (app *application) getBookDetailHandler(w http.ResponseWriter, r *http.Request) {
    db, err := sql.Open(dbDriver, os.Getenv("DATABASE_URL"))
    if err != nil {
      panic(err.Error())
    }
    defer db.Close()
    
    idStr := r.URL.Query().Get("id")

    bookID, err := strconv.Atoi(idStr)

    book, err := getBookDetail(app, db, bookID)
    if err != nil {
      app.logger.Error(err.Error())
      http.Error(w, "Book not found", http.StatusNotFound)
      return
    }
   
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(book)
}

func getBookDetail(app *application, db *sql.DB, id int) (*BookDetail, error) {
    query := "SELECT * FROM books WHERE id = $1;"
    row := db.QueryRow(query, id)

    book := &BookDetail{}
    err := row.Scan(&book.Id, &book.Title, &book.Author, &book.Description, &book.Rating)
    if err != nil {    	
        return nil, err
    }
    return book, nil
}

func (app *application) deleteBookHandler(w http.ResponseWriter, r *http.Request) {
    db, err := sql.Open(dbDriver, os.Getenv("DATABASE_URL"))
    if err != nil {
      panic(err.Error())
    }
    defer db.Close()
    
    idStr := r.URL.Query().Get("id")

    bookID, err := strconv.Atoi(idStr)

    err = deleteBook(app, db, bookID)
    if err != nil {
      app.logger.Error(err.Error())
      http.Error(w, "delete failed", http.StatusNotFound)
      return
    }
    w.WriteHeader(http.StatusOK)
}

func deleteBook(app *application, db *sql.DB, id int) (error) {
    query := "DELETE FROM books WHERE id = $1;"
    _, err := db.Exec(query, id)

    return err
}

func (app *application) updateBookHandler(w http.ResponseWriter, r *http.Request) {
    db, err := sql.Open(dbDriver, os.Getenv("DATABASE_URL"))
    if err != nil {
      panic(err.Error())
    }
    defer db.Close()
    
    var book BookDetail
    json.NewDecoder(r.Body).Decode(&book)

    err = updateBook(app, db, book)
    if err != nil {
      app.logger.Error(err.Error())
      http.Error(w, "update failed", http.StatusNotFound)
      return
    }  
    w.WriteHeader(http.StatusOK)
}

func updateBook(app *application, db *sql.DB, book BookDetail) (error) {
    query := "UPDATE books SET title = '" + book.Title + "'" +
    		", author = '" + book.Author + "'" +
    		", description = '" + book.Description + "'" +
    		", rating = " + strconv.Itoa(book.Rating) +
    		" WHERE id = $1;"
    _, err := db.Exec(query, book.Id)

    return err
}

func (app *application) createBookHandler(w http.ResponseWriter, r *http.Request) {
    db, err := sql.Open(dbDriver, os.Getenv("DATABASE_URL"))
    if err != nil {
      panic(err.Error())
    }
    defer db.Close()
    
    var book BookDetail
    json.NewDecoder(r.Body).Decode(&book)

    err = createBook(app, db, book)
    if err != nil {
      app.logger.Error(err.Error())
      http.Error(w, "create failed", http.StatusNotFound)
      return
    }  
    w.WriteHeader(http.StatusCreated)
}

func createBook(app *application, db *sql.DB, book BookDetail) (error) {
    query := "INSERT INTO books(title, author, description, rating) VALUES (" +
    		"'" + book.Title + "'," +
    		"'" + book.Author + "'," +
    		"'" + book.Description + "'," +
    		"'" + strconv.Itoa(book.Rating) + "'" +
    		");"
    _, err := db.Exec(query)

    return err
}
