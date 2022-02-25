package handlers

import (
	"fmt"
	"net/http"
)

func WelcomePage(writer http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprintf(writer, "welcome to dexters first page")
	fmt.Println("detected call to welcome-page")
}