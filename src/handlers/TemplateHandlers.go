package handlers

import (
  "html/template"
  "log"
  "net/http"
  "path"
)


func ServeTemplate(w http.ResponseWriter, r *http.Request) {
  log.Print("I ask index")
  lp := path.Join("../web", "index.html")
  tmpl, err := template.ParseFiles(lp)
  if err != nil {
  	log.Println(err)
  }
  tmpl.ExecuteTemplate(w, "index.html", nil)
}

func ServeStyle(w http.ResponseWriter, r *http.Request) {
  log.Print("I ask style")
  lp := path.Join("../web/style/", "style.css")
  tmpl, err := template.ParseFiles(lp)
  if err != nil {
    log.Println(err)
  }
  //style : template.NewCSS(lp)
  error := tmpl.Execute(w, "")
  if error != nil {
    log.Println(error)
  } 
}

func ServeJs(w http.ResponseWriter, r *http.Request) {
  log.Print("I ask Js")
  lp := path.Join("../web/js/", "AJAXform.js")
  tmpl, err := template.ParseFiles(lp)
  if err != nil {
    log.Println(err)
  }
  //style : template.NewCSS(lp)
  error := tmpl.Execute(w, "")
  if error != nil {
    log.Println(error)
  } 
}   