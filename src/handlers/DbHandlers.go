package handlers

import (
  "../actions"
  "log"
  "net/http"
  "../modele"
  "encoding/json"
)


func GetUser(w http.ResponseWriter, r *http.Request) {
  //vars := mux.Vars(r)
} 


func AddUser(w http.ResponseWriter, r *http.Request) {
  decoder := json.NewDecoder(r.Body)
  data := modele.Formulaire{}
  error := decoder.Decode(&data)
  if error != nil {
    log.Print(error)
  }
  errorAddUser := actions.AddUser(data)
  if errorAddUser != nil {
    log.Print(errorAddUser)
  }
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

}   
 