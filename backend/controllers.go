package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type Response struct {
	Message string
	StatusCode int
}

type UserInput struct {
	Email string
	Username string
	Password string
}

func CreateUserController(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(writer, "Wrong method", http.StatusMethodNotAllowed)
		return
	}
	if req.ContentLength == 0 {
		http.Error(writer, "No data sent", http.StatusBadRequest)
		return
	}


	data, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(writer, "Server error", http.StatusInternalServerError)
		return
	}

	var newUser UserInput
	err = json.Unmarshal(data, &newUser)
	if err != nil {
		http.Error(writer, "Server error", http.StatusInternalServerError)
		return
	}

	// Acá debería agregar validaciones. No sé si meter una librería o hacerlo a mano

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
	newUser.Password = string(hashedPassword)

	if err != nil {
		return
	}

	err = CreateUser(newUser)
	if err != nil {
		http.Error(writer, "Server error", http.StatusInternalServerError)
		return
	}
	
	response := Response{ Message: "User created succesfully", StatusCode: 201}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(writer, "Server error", http.StatusInternalServerError)
		return
	}

	writer.Write(jsonResponse)
}

type ReadUserInput struct {
	Id int
}

func ReadUserController(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(writer, "Bad request", http.StatusBadRequest)
		return
	}

	if req.ContentLength == 0 {
		http.Error(writer, "No data sent", http.StatusBadRequest)
		return
	}

	data, err := io.ReadAll(req.Body)
	if err != nil {
		log.Fatal("Error: ", err)
		http.Error(writer, "Server error", http.StatusInternalServerError)
		return
	}

	var userId ReadUserInput
	err = json.Unmarshal(data, &userId)
	if err != nil {
		log.Fatal("Error: ", err)
		http.Error(writer, "Server error", http.StatusInternalServerError)
		return
	}


	user, err := ReadUser(userId.Id)
	if err != nil {
		log.Fatal("Error: ", err)
		http.Error(writer, "Server error", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		http.Error(writer, "Server error", http.StatusInternalServerError)
		return
	}

	writer.Write(response)
}

func UpdateUserController(writer http.ResponseWriter, req *http.Request) {
}

func DeleteUserController(writer http.ResponseWriter, req *http.Request) {
}
