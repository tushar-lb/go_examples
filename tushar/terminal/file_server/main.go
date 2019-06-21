package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/tushar/terminal/file_server/server"
)

func main() {
	port := flag.String("port", "8000", "port to start rest server on")
	isHttps := flag.Bool("HTTPS", false, "is https")
	flag.Parse()

	fmt.Printf("Bool %v", *isHttps)
	router := server.NewRouter()
	log.Println("starting rest file server")
	log.Fatal(http.ListenAndServe("0.0.0.0:"+*port, router))
}
