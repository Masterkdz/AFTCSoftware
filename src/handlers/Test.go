package handlers

import (
	"net/http"
	"log"
)

func Test (w http.ResponseWriter, r *http.Request) {
	log.Println("OH OH AH AH")
}

func TestDB (w http.ResponseWriter, r *http.Request) {
	log.Println("TESTDBROUTE")
}