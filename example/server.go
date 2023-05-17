package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	var port string
	flag.StringVar(&port, "p", "8080", "port to listen on")
	flag.Parse()

	mux := http.NewServeMux()
	mux.Handle("/api/test", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		defer r.Body.Close()
		fmt.Println("BODY:", string(b))
		fmt.Println("HEADERS:", r.Header)
		fmt.Println("METHOD:", r.Method)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}))
	server := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	fmt.Println("Listening on port", port)

	server.ListenAndServe()
}
