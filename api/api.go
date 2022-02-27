package api

import (
	"autogate/api/handlers"
	"autogate/api/middlewares"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ojoadeolagabriel2/autogate-go-core/util"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

const (
	DefaultPort   = 12345
	EnvAppPortKey = "ENV_APP_PORT"
)

var EnvApiPrefixKey = util.GetConfig("ENV_API_VERSION_PREFIX_KEY", "/v1")

func getAppPort() int {
	 return util.GetIntConfig(EnvAppPortKey, DefaultPort)
}

func RegisterRoutes(router *mux.Router) {
	fmt.Println("Autogate Rest API - v1.0")

	// defaults
	router.Use(middlewares.CommonMiddleware)
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}).Methods("GET")

	// user handler mapping
	router.HandleFunc(EnvApiPrefixKey + "/users", handlers.GetAllUsers).Methods("GET")
	router.HandleFunc(EnvApiPrefixKey + "/users", handlers.CreateUser).Methods("POST")
	router.HandleFunc(EnvApiPrefixKey + "/users/{id}", handlers.GetUserById).Methods("GET")

	// transaction handler mapping
	router.HandleFunc(EnvApiPrefixKey + "/transaction/{id}", handlers.GetTransferById).Methods("GET")
	router.HandleFunc(EnvApiPrefixKey + "/transaction/{id}/qrcode", handlers.GetQrCode).Methods("GET")

	// handle errors
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(fmt.Sprintf(":%d", getAppPort()), router)
	if err != nil {
		log.Fatalln("test failed: " + err.Error())
	}
}
