package modele

type Formulaire struct {
    Date		string
    Nomcontact	string 
    Nomtc       string
    Adresse		string  
    Tel			string
    Auteurfiche	string
    Lien		string
    Circonstances	string 
    Priseenchargemedicale	string 
    Priseenchargesociale	string
    Demarches	string
    Problemes	string
    Attentes	string
    Observations	string
}

type Adherent struct {
    Nom string
    Prenom string
    Date string
    Tel string
}

type Adherents []Adherent