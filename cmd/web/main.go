package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// application struct to handle application-wide dependencies
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// command line argument
	addr := flag.String("addr", ":8000", "HTTP Network Address")
	flag.Parse()

	// Define new loggers for information messages and error messages
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/create", app.createSnippet)
	mux.HandleFunc("/snippet", app.showSnippet)

	// File server to serve static files
	fileServer := http.FileServer(http.Dir("../../ui/static/"))

	// Register filesServer as handler
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Handler http errors with our custom loggers
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Println("Starting server on %s", *addr)

	// Call ListenAndServe() method from custom http.Server struct
	err := srv.ListenAndServe()
	errorLog.Fatal(err)

}
