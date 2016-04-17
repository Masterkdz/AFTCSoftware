/*package main 

import (
	"./handlers"
	"log"
	"net/http"
	"html/template"
	"./router"
	"fmt"
)

func handler(w http.ResponseWriter, r *http.Request) {
   fmt.Fprintf(w, "test", r.URL.Path[1:])
}

func main() {
	router := router.NewRouter()
	fs := http.FileServer(http.Dir("template"))
  	http.Handle("/", fs)
  	http.HandleFunc("/handler", test.Test)

	log.Println("API listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))router))
}*/

package main

import (
  "html/template"
  "log"
  "net/http"
  "path"
)

func main() {
  //fs := http.FileServer(http.Dir("template"))
  //http.Handle("./template/", http.StripPrefix("./template/", fs))

  http.HandleFunc("/", serveTemplate)

  log.Println("Listening...")
  http.ListenAndServe(":8080", nil)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
  lp := path.Join("./template", "index.html")
  //fp := path.Join("	./template", r.URL.Path)

  tmpl, err := template.ParseFiles(lp)
  if err != nil {
  	log.Println(err)
  }
  tmpl.ExecuteTemplate(w, "index.html", nil)
}