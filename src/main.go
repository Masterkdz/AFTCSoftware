package main 

import (
	//"./handlers"
	"log"
	"net/http"
	//"html/template"
	//"./router"
	//"fmt"
)

//func handler(w http.ResponseWriter, r *http.Request) {
//    fmt.Fprintf(w, "test", r.URL.Path[1:])
//}

func main() {
	//router := router.NewRouter()
	fs := http.FileServer(http.Dir("template"))
  	http.Handle("/", fs)
  	//http.HandleFunc("/handler", test.Test)

	log.Println("API listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))//router))
}