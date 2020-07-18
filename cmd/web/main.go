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
	// Parsing runtime config settings
	addr := flag.String("addr", ":8000", "HTTP Network Address")
	flag.Parse()

	// Define new loggers for information messages and error messages
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// Custom http.Server struct
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Println("Starting server on %s", *addr)

	// Call ListenAndServe() method from custom http.Server struct
	err := srv.ListenAndServe()
	errorLog.Fatal(err)

}
