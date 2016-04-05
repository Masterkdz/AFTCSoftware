package actions

import (
	"../../poller"
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"log"
	_ "github.com/go-sql-driver/mysql"
    "fmt"
)


func AddResult(db *sql.DB,) ([]string, error) {
	var result []string
	db, error := sql.Open("mysql", "Root@/db.sql")
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
}

func TestDispResult(db *sql.DB,) ([]string, error) {
	var result []string
	db, err := sql.Open("mysql", "Root@/db.sql")
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