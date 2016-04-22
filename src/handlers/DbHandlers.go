package handlers

import (
  //"../actions"
  "../utils"
  "log"
  "net/http"
  "../modele"
  "encoding/json"
)


func GetUser(w http.ResponseWriter, r *http.Request) {
  //vars := mux.Vars(r)
} 


func AddUser(w http.ResponseWriter, r *http.Request) {
  header := w.Header()
  utils.SetCors(&header)
  decoder := json.NewDecoder(r.Body)
  log.Print(decoder)
  var dat modele.Formulaire
  error := decoder.Decode(&dat)
  if error != nil {
    log.Print("erreor decode ", error)
  }

  log.Print(dat)
/*  errorAddUser := actions.AddUser(data)
  if errorAddUser != nil {
    log.Print(errorAddUser)
  }*/
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

}   
 