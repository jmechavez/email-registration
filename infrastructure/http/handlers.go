package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jmechavez/email-registration/internal/ports/service"
)

type UserHandlers struct {
	service service.UserService
}

func (uh *UserHandlers) GetAllUsersEmail(w http.ResponseWriter, r *http.Request) {
	users, err := uh.service.GetAllUserEmails()
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, users)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(fmt.Sprintf("Failed to encode response %v", err))
	}
}
