package main 

import (
	"log"
	"net/http"
	"./router"
)

func main() {
	router := router.NewRouter()
	log.Println("API listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}