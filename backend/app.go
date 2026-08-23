package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main () {
	err := InitializeDB()
	if err != nil {
		log.Fatal("Error: ", err)
		return
	}

	http.HandleFunc("/users/create", CreateUserController)
	http.HandleFunc("/users/read", ReadUserController)

	fmt.Println("Servidor levantado en puerto 9000")
	err = http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("Error: ", err)
		return
	}
}
