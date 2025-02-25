package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jmechavez/email-registration/internal/dto"
	"github.com/jmechavez/email-registration/internal/ports/service"
)

type UserHandlers struct {
	service service.UserService
}

type DelUserHandlers struct {
	service service.DelUserService
}

func (uh *UserHandlers) GetAllUsersEmail(w http.ResponseWriter, r *http.Request) {
	users, err := uh.service.GetAllUserEmails()
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, users)
	}
}

func (uh *UserHandlers) NewUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idNo := vars["id_no"]
	var request dto.NewUserEmailRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.IdNo = idNo
		user, appError := uh.service.NewUser(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.Message)
		} else {
			writeResponse(w, http.StatusCreated, user)
		}
	}
}

func (uh *DelUserHandlers) DelUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idNo := vars["id_no"]
	var request dto.DeleteUserEmailRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	request.IdNo = idNo
	user, appError := uh.service.DelUser(request)
	if appError != nil {
		writeResponse(w, appError.Code, appError.Message)
		return
	}
	writeResponse(w, http.StatusCreated, user)
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow frontend
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(code)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(fmt.Sprintf("Failed to encode response %v", err))
	}
}
