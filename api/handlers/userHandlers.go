package handlers

import (
	"autogate/data"
	"autogate/util"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

const (
	IdKey = "id"
)

func GetUserById(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	key := vars[IdKey]

	for _, user := range data.Users {
		val, _ := strconv.ParseInt(key, 10, 64)
		if user.Id == val {
			payload, err := json.Marshal(user)
			if err != nil {
				fmt.Println(err.Error())
			}
			writer.Header().Set("Content-Type", "application/json")
			_, _ = writer.Write(payload)
			return
		}
	}

	emptyResponse, _ := json.Marshal(make(map[string]string))
	_, err := writer.Write(emptyResponse)
	if err != nil {
		log.Println(err.Error())
	}
}

func GetAllUsers(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	payload, _ := json.Marshal(data.Users)
	_, _ = writer.Write(payload)
}

func CreateUser(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			payload := apiError(util.UnknownError, "could not complete")
			_, _ = fmt.Fprintf(writer, payload)
		}
	}()

	var user data.User
	err := json.NewDecoder(request.Body).Decode(&user)
	validateUser(&user)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		payload := apiError(util.UnknownError, err.Error())
		_, _ = fmt.Fprintf(writer, payload)
		return
	}

	users := append(data.Users, user)
	data.Users = users

	if payload, err := json.Marshal(users); err == nil {
		writer.Header().Set("Content-Type", "application/json")
		_, _ = writer.Write(payload)
	}
}

func apiError(code string, message string) string {
	return fmt.Sprintf("{\"error_code\": \"%s\", \"error_message\": \"%s\"}", code, message)
}

func validateUser(user *data.User) {
	if user == nil ||
		user.Id <= 0 {
		panic("invalid user supplied")
	}
}
