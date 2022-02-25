package handlers

import (
	"autogate/data"
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

func HandleGetAllUser(writer http.ResponseWriter, request *http.Request) {
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

func HandleGetAllUsers(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	payload, _ := json.Marshal(data.Users)
	_, _ = writer.Write(payload)
}
