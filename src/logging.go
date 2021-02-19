package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type logging struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func logout(writer http.ResponseWriter, req *http.Request) {
	sessions.DeleteSession(&writer)
}

func login(writer http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var loginData logging

	err := json.NewDecoder(req.Body).Decode(&loginData)
	fmt.Println(loginData)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if loginData.Password == "admin" {
		userSession := sessions.NewSession(&writer)
		userSession.UserId = mongodb.GetOrCreateUser(loginData.Email)
	} else {
		writer.WriteHeader(http.StatusBadRequest)
	}
}
