package main

import (
	"net/http"
)

func HomeController(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(writer, "Wrong method", http.StatusMethodNotAllowed)
		return
	}

	writer.Write([]byte("Bienvenido a la API. Haz cualquier request :)"))
}

func CreateUserController(writer http.ResponseWriter, req *http.Request) {
}

func ReadUserController(writer http.ResponseWriter, req *http.Request) {
}

func UpdateUserController(writer http.ResponseWriter, req *http.Request) {
}

func DeleteUserController(writer http.ResponseWriter, req *http.Request) {
}
