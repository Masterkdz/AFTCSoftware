package handlers

import (
  "html/template"
  "log"
  "net/http"
  "path"
)

func Test (w http.ResponseWriter, r *http.Request) {
	log.Println("test")
}

func TestDB (w http.ResponseWriter, r *http.Request) {
	log.Println("TESTDBROUTE")
}

func ServeTemplate(w http.ResponseWriter, r *http.Request) {
  lp := path.Join("./template", "index.html")
  //cssHandler := http.FileServer(http.Dir("./template/style"))
  //http.Handle("./template/style", http.StripPrefix("./template/style", cssHandler))
  //fp := path.Join("	./template", r.URL.Path)

  tmpl, err := template.ParseFiles(lp)
  if err != nil {
  	log.Println(err)
  }
  tmpl.ExecuteTemplate(w, "index.html", nil)
}