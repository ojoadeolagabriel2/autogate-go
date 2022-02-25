package api

import (
	"autogate/api/handlers"
	"autogate/api/middlewares"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strconv"
)

const (
	DefaultPort = 12345
	EnvAppPortKey = "ENV_APP_PORT"
)

func getAppPort() int64 {
	port, err := strconv.ParseInt(os.Getenv(EnvAppPortKey), 10, 64)
	if err != nil {
		fmt.Printf("%s not set, defaulting to %d", EnvAppPortKey, DefaultPort)
		return DefaultPort
	}
	return port
}

func HandleRequests(router *mux.Router) {
	fmt.Println("Autogate Rest API - v1.0")
	router.Use(middlewares.CommonMiddleware)

	// welcome handler mapping
	router.HandleFunc("/", handlers.WelcomePage)

	// user handler mapping
	router.HandleFunc("/users", handlers.HandleGetAllUsers)
	router.HandleFunc("/users/{id}", handlers.HandleGetAllUser)

	// transaction handler mapping
	router.HandleFunc("/transaction/{id}", handlers.HandleGetTransferById)

	// handle errors
	go http.ListenAndServe(fmt.Sprintf(":%d", getAppPort()), router)
}
