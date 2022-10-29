package main

import (
	"context"
	"flag"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	pool, err := pgxpool.New(context.Background(), os.Getenv("postgres://postgres:7894561230@localhost:5432/snippetbox"))
	if err != nil {
		errorLog.Fatalf("Unable to connection to database: %v\n", err)
	}
	defer pool.Close()
	infoLog.Print("Connected!")

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,

		Handler: app.routes(),
	}
	infoLog.Printf("Starting server on %s", *addr)
	infoLog.Println("http://localhost:4000")
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

//func main() {
//	addr := flag.String("addr", ":4000", "HTTP network address")
//	flag.Parse()
//	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
//	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
//	app := &application{
//		errorLog: errorLog,
//		infoLog:  infoLog,
//	}
//	srv := &http.Server{
//		Addr:     *addr,
//		ErrorLog: errorLog,
//		// Call the new app.routes() method to get the servemux containing our routes.
//		Handler: app.routes(),
//	}
//	infoLog.Printf("Starting server on %s", *addr)
//	err := srv.ListenAndServe()
//	errorLog.Fatal(err)
//}

//func main() {
//	addr := flag.String("addr", ":4000", "HTTP network address")
//	flag.Parse()
//	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
//	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
//	// Initialize a new instance of our application struct, containing the
//	// dependencies.
//	app := &application{
//		errorLog: errorLog,
//		infoLog:  infoLog,
//	}
//	// Swap the route declarations to use the application struct's methods as the
//	// handler functions.
//	mux := http.NewServeMux()
//	fileServer := http.FileServer(http.Dir("./ui/static/"))
//	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
//	mux.HandleFunc("/", app.home)
//	mux.HandleFunc("/snippet/view", app.snippetView)
//	mux.HandleFunc("/snippet/create", app.snippetCreate)
//	srv := &http.Server{
//		Addr:     *addr,
//		ErrorLog: errorLog,
//		Handler:  mux,
//	}
//	infoLog.Printf("Starting server on %s", *addr)
//	err := srv.ListenAndServe()
//	errorLog.Fatal(err)
//}
