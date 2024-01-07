package main

import (
	"embed"
	"flag"
	"log"
	"net/http"
	"os"
)

var logger = log.New(os.Stdout, "[INFO]\t", log.LstdFlags)

//go:embed "*"
var Files embed.FS

func main() {

	addr := flag.String("addr", ":9000", "The addr of the application.")
	flag.Parse()

	logger.Println("Starting web server on", *addr)

	server := &http.Server{
		Addr:    *addr,
		Handler: routes(),
	}

	err := server.ListenAndServeTLS("./server.crt", "./server.key")

	logger.Fatal(err)
}

func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Printf("%s %s %s", r.Proto, r.Method, r.URL.RequestURI())
		next.ServeHTTP(w, r)
	})
}

func routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.FS(Files))
	mux.Handle("/", fileServer)
	return logRequest(mux)
}
