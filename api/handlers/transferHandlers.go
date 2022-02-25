package handlers

import (
	"net/http"
)

func HandleGetTransferById(writer http.ResponseWriter, request *http.Request) {
	_, _ = writer.Write([]byte("<in development>"))
}