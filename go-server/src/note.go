package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type editNote struct {
	Note string `json:"note"`
}

func note(writer http.ResponseWriter, req *http.Request) {
	userSession := sessions.GetSession(req)
	if userSession == nil {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	if req.Method == http.MethodGet {
		fmt.Fprint(writer, mongodb.GetUser(userSession.UserId).NOTE)
		return
	}

	if req.Method == http.MethodPost {
		var editNote editNote
		err := json.NewDecoder(req.Body).Decode(&editNote)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}
		mongodb.UpdateUserNote(userSession.UserId, editNote.Note)
		return
	}

	writer.WriteHeader(http.StatusMethodNotAllowed)
}
