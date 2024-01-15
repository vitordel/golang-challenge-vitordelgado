package main

import (
	"database/sql"
	"fmt"
	"golang-challenge-vitordelgado/routes"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB

func main() {
	var err error
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/flexconsulta")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	routes.SetupRoutes(router)

	http.Handle("/", router)
	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
