package handlers

import (
  "../utils"
  "html/template"
  "log"
  "net/http"
  "path"
)


func ServeTemplate(w http.ResponseWriter, r *http.Request) {
  header := w.Header()
  utils.SetCors(&header)
  log.Print("I ask index")
  lp := path.Join("../web", "index.html")
  tmpl, err := template.ParseFiles(lp)
  if err != nil {
  	log.Println(err)
  }
  tmpl.ExecuteTemplate(w, "index.html", nil)
}

func ServeStyle(w http.ResponseWriter, r *http.Request) {
  header := w.Header()
  utils.SetCors(&header)
  log.Print("I ask style")
  w.Header().Set("Content-type", "text/css")
  http.ServeFile(w, r, "../web/style/style.css")
}

func ServeJs(w http.ResponseWriter, r *http.Request) {
  header := w.Header()
  utils.SetCors(&header)
  log.Print("I ask AjaxJs")
  lp := path.Join("../web/js/", "AJAXform.js")
  w.Header().Set("Content-type", "application/javascript")
  http.ServeFile(w, r, lp)
}   

func ServeJQuery(w http.ResponseWriter, r *http.Request) {
  header := w.Header()
  utils.SetCors(&header)
  log.Print("I ask JQuery")
  w.Header().Set("Content-type", "application/javascript")
  http.ServeFile(w, r, "../web/js/jquery1.12.3.js")
}   