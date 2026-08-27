package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)


func UserController(writer http.ResponseWriter, req *http.Request) {
	var jsonResponse []byte

	switch req.Method {

	case http.MethodGet:
		pathValue := req.PathValue("id")
		if pathValue == "" {
			http.Error(writer, "Bad request", http.StatusBadRequest)
			return
		}
		userId, err := strconv.ParseInt(pathValue, 10, 0)
		if err != nil {
			log.Fatal("Error: ", err)
			http.Error(writer, "Server error", http.StatusInternalServerError)
			return
		}

		user, err := ReadUser(userId)
		if err != nil {
			log.Fatal("Error: ", err)
			http.Error(writer, "Server error", http.StatusInternalServerError)
			return
		}

		jsonResponse, err = json.Marshal(user)
		if err != nil {
			http.Error(writer, "Server error", http.StatusInternalServerError)
			return
		}

	case http.MethodPost:
		if req.ContentLength == 0 {
			log.Fatal("Error: no hay contenido")
			http.Error(writer, "No data sent", http.StatusBadRequest)
			return
		}

		data, err := io.ReadAll(req.Body)
		if err != nil {
			log.Fatal("Error: ", err)
			http.Error(writer, "Server error", http.StatusInternalServerError)
			return
		}

		var newUser NewUserInput
		err = json.Unmarshal(data, &newUser)
		if err != nil {
			log.Fatal("Error: ", err)
			http.Error(writer, "Server error", http.StatusInternalServerError)
			return
		}

		// Acá debería agregar validaciones. No sé si meter una librería o hacerlo a mano

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
		if err != nil {
			log.Fatal("Error: ", err)
			http.Error(writer, "Server error", http.StatusInternalServerError)
			return
		}
		newUser.Password = string(hashedPassword)


		createdUser, err := CreateUser(newUser)
		if err != nil {
			log.Fatal("Error: ", err)
			http.Error(writer, "Server error", http.StatusInternalServerError)
			return
		}

		jsonResponse, err = json.Marshal(createdUser)
		if err != nil {
			log.Fatal("Error: ", err)
			http.Error(writer, "Server error", http.StatusInternalServerError)
			return
		}

		fmt.Println("Usuario creado")
	case http.MethodPut:
		return
	case http.MethodDelete:
		return
	default:
		fmt.Println("No se esperaba ese método")
	}

	writer.Write(jsonResponse)
}
