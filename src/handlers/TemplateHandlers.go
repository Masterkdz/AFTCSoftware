package handlers

import (
  "html/template"
  "log"
  "net/http"
  "path"
)


func ServeTemplate(w http.ResponseWriter, r *http.Request) {
  lp := path.Join("./template", "index.html")
  tmpl, err := template.ParseFiles(lp)
  if err != nil {
  	log.Println(err)
  }
  tmpl.ExecuteTemplate(w, "index.html", nil)
}

func ServeStyle(w http.ResponseWriter, r *http.Request) {
  lp := path.Join("./template/style/", "style.css")
  log.Print(lp)
  http.FileServer(http.Dir(lp))
}   