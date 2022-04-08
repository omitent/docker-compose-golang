package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Test struct {
	nombre string
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", DoHealthCheck).Methods("GET")
	router.HandleFunc("/test", SendTest).Methods("GET")
	router.HandleFunc("/version", GetVersion).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func DoHealthCheck(w http.ResponseWriter, r *http.Request) {
	names := []string{}
	// Open up our database connection.
	db, err := sql.Open("mysql", "root:hhx999@tcp(db)/db_app")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT * FROM test")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		test := new(Test)
		err := results.Scan(&test.nombre)
		if err != nil {
			panic(err.Error())
		}
		names = append(names, test.nombre)
	}
	res := strings.Join(names, ",")
	log.Println(res)
	fmt.Fprintf(w, res)
}

func GetVersion(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:hhx999(db)/db_app")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	var version string

	err2 := db.QueryRow("SELECT VERSION()").Scan(&version)

	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Fprintf(w, version)
}

func SendTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Imprimiendo")
}
