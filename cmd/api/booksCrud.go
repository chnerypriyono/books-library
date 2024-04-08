package main

import (
	"net/http"
	"database/sql"
    "encoding/json"       
    "strconv"
    _ "github.com/lib/pq"   
    "os"
    "context"
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

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    _, err := app.verifyIDToken(ctx, r)
    if err != nil {
        http.Error(w, "Authorization Required", http.StatusUnauthorized)
        return
    }

    db, err := sql.Open(dbDriver, os.Getenv("DATABASE_URL"))
    if err != nil {
      panic(err.Error())
    }
    defer db.Close()    

    books, err := getBooks(app, db)
    if err != nil {
      app.logger.Error(err.Error())
      http.Error(w, "Get Books Failed", http.StatusInternalServerError)
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
        books = append(books, book)
    }
    if err = rows.Err(); err != nil {    	
        return books, err
    }
    return books, nil
}

func (app *application) deleteBookHandler(w http.ResponseWriter, r *http.Request) {
    
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    _, err := app.verifyIDToken(ctx, r)
    if err != nil {
        http.Error(w, "Authorization Required", http.StatusUnauthorized)
        return
    }

    db, err := sql.Open(dbDriver, os.Getenv("DATABASE_URL"))
    if err != nil {
      panic(err.Error())
    }
    defer db.Close()
    
    idStr := r.URL.Query().Get("id")

    bookID, err := strconv.Atoi(idStr)
    if err != nil {
      app.logger.Error(err.Error())
      http.Error(w, "Invalid 'id' query param value for deleteBook", http.StatusBadRequest)
      return
    }

    err = deleteBook(app, db, bookID)
    if err != nil {
      app.logger.Error(err.Error())
      http.Error(w, "Delete Book Failed", http.StatusInternalServerError)
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

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    _, err := app.verifyIDToken(ctx, r)
    if err != nil {
        http.Error(w, "Authorization Required", http.StatusUnauthorized)
        return
    }

    db, err := sql.Open(dbDriver, os.Getenv("DATABASE_URL"))
    if err != nil {
      panic(err.Error())
    }
    defer db.Close()
    
    var book Book
    err = json.NewDecoder(r.Body).Decode(&book)
    if err != nil {
      app.logger.Error(err.Error())
      http.Error(w, "Invalid request body", http.StatusBadRequest)
      return
    }

    err = updateBook(app, db, book)
    if err != nil {
      app.logger.Error(err.Error())
      http.Error(w, "Update Book Failed", http.StatusInternalServerError)
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

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    _, err := app.verifyIDToken(ctx, r)
    if err != nil {
        http.Error(w, "Authorization Required", http.StatusUnauthorized)
        return
    }

    db, err := sql.Open(dbDriver, os.Getenv("DATABASE_URL"))
    if err != nil {
      panic(err.Error())
    }
    defer db.Close()
    
    var book Book
    err = json.NewDecoder(r.Body).Decode(&book)
    if err != nil {
      app.logger.Error(err.Error())
      http.Error(w, "Invalid request body", http.StatusBadRequest)
      return
    }

    err = createBook(app, db, book)
    if err != nil {
      app.logger.Error(err.Error())
      http.Error(w, "Create Book Failed", http.StatusInternalServerError)
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
