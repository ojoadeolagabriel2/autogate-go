package handlers

import (
	"autogate/data"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
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
			payload, err := json.Marshal(data.Users)
			if err != nil {
				fmt.Println(err.Error())
			}
			writer.Header().Set("Content-Type", "application/json")
			_, _ = writer.Write(payload)
			break
		}
	}
}

func HandleGetAllUsers(writer http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprintf(writer, "welcome to dexters first page")
	fmt.Println("detected call to welcome-page")
}
