package actions

import (
	"../modele"
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func AddUser(data modele.Formulaire) error {
	db, error := sql.Open("mysql", "Root@/AFTC")
    if error != nil {
    	log.Print(error)
    }
	if db == nil {
		error := errors.New("db = nil ")
		return error
	}   
	stmt, _ := db.Prepare("INSERT INTO usagers(Date, Nomcontact, Nomtc, Adresse, Tel, Auteurfiche, Lien, Circonstances, Priseenchargemedicale, Priseenchargesociale, Demarches, Problemes, Attentes, Observations) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")

    res, errExec := stmt.Exec(data.Date, data.Nomcontact, data.Nomtc, data.Adresse, data.Tel, data.Auteurfiche, data.Lien, data.Circonstances, data.Priseenchargemedicale, data.Priseenchargesociale, data.Demarches, data.Problemes, data.Attentes, data.Observations) //Remplacer les valeurs par les variables corresepondantes
    if errExec != nil {
    	log.Print(errExec)
    }
    _, _ = res.LastInsertId()
	db.Close()
	return nil
}

func GetUsers(data *modele.Adherents) ([]string, error) {
	var result []string
	db, err := sql.Open("mysql", "Root@/dbAFTC.db3")
    checkErr(error)
	if db == nil {
		err := errors.New("db = nil ")
		return result, err
	}
	rows, err := db.Query("SELECT nom, prenom FROM usagers")
    checkErr(err)

    for rows.Next() {
        var nom string
        var prenom string
        err = rows.Scan(&nom, &prenom)
        checkErr(err)
        fmt.Println(nom)
        fmt.Println(prenom)
    }

	if err != nil {
        panic(err)
    }
    
	db.Close()
	return result, nil
}

/*func TestAddResult(db *sql.DB,) ([]string, error) {
	var result []string
	db, error := sql.Open("mysql", "Root@/dbAFTC.db3")
    checkErr(error)
	if db == nil {
		error := errors.New("db = nil ")
		return result, error
	}
	stmt, err := db.Prepare("INSERT INTO usagers(nom, prenom) VALUES(?,?)")
    checkErr(err)

    res, err := stmt.Exec("DEZORZI", "Kevin")
    checkErr(err)

    id, err := res.LastInsertId()
    checkErr(err)
	if errorQuery != nil {
		return result, errorQuery
	}
	db.Close()
	return result, nil
}*/



/*func DisplayResult(db *sql.DB,) ([]string, error) {
	var result []string
	db, err := sql.Open("mysql", "Root@/dbAFTC.db3")
    checkErr(error)
	if db == nil {
		err := errors.New("db = nil ")
		return result, err
	}
	rows, err := db.Query("SELECT nom, prenom, Auteur_fiche, Date, NomTCCL, telephone_fixe, Portable, CirconstancesTrauma, ConditionsPriseEnchargeMedicale, ConditionsPriseEnChargeSociale, DemarchesAdministrativesJuridiquesEffectuees, ProblemesRecontres, AttentesAFTC, Observations FROM usagers")
    checkErr(err)

    for rows.Next() {
        var nom string
        var prenom string
        err = rows.Scan(&nom, &prenom, &Auteur_fiche, &Date, &NomTCCL, &telephone_fixe, &Portable, &CirconstancesTrauma, &ConditionsPriseEnchargeMedicale, aConditionsPriseEnChargeSociale, &DemarchesAdministrativesJuridiquesEffectuees, &ProblemesRecontres, &AttentesAFTC, &Observations)
        checkErr(err)
        fmt.Println(nom)
        fmt.Println(prenom)
    }

	if err != nil {
        panic(err)
    }
    
	db.Close()
	return result, nil
}*/
    