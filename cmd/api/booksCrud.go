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

type Book struct {
	Id    		int 	`json:"id"`
    Title  		string	`json:"title"`
    Author 		string	`json:"author"`
    Publisher   string  `json:"publisher"`
	Description string	`json:"description"`
	ImageUrl    string  `json:"image_url"`
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

func getBooks(app *application, db *sql.DB) ([]Book, error) {
    query := "SELECT id, title, author, publisher, description, imageurl FROM books ORDER BY title ASC;"
    rows, err := db.Query(query)

    if err != nil {
        return nil, err
    }
    defer rows.Close()

    books := []Book{}

    for rows.Next() {
        var book Book
        if err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.Publisher, &book.Description, &book.ImageUrl); err != nil {
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
    
    var book Book
    json.NewDecoder(r.Body).Decode(&book)

    err = updateBook(app, db, book)
    if err != nil {
      app.logger.Error(err.Error())
      http.Error(w, "update failed", http.StatusNotFound)
      return
    }  
    w.WriteHeader(http.StatusOK)
}

func updateBook(app *application, db *sql.DB, book Book) (error) {
    query := "UPDATE books SET title = '" + book.Title + "'" +
    		", author = '" + book.Author + "'" +
            ", publisher = '" + book.Publisher + "'" +
    		", description = '" + book.Description + "'" +
    		", imageurl = '" + book.ImageUrl + "'" +
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
    
    var book Book
    json.NewDecoder(r.Body).Decode(&book)

    err = createBook(app, db, book)
    if err != nil {
      app.logger.Error(err.Error())
      http.Error(w, "create failed", http.StatusNotFound)
      return
    }  
    w.WriteHeader(http.StatusCreated)
}

func createBook(app *application, db *sql.DB, book Book) (error) {
    query := "INSERT INTO books(title, author, publisher, description, imageurl) VALUES (" +
    		"'" + book.Title + "'," +
    		"'" + book.Author + "'," +
            "'" + book.Publisher + "'," +
    		"'" + book.Description + "'," +
    		"'" + book.ImageUrl + "'" +
    		");"
    _, err := db.Exec(query)

    return err
}
