package handlers

import (
	"log"
	"net/http"
)

func HandleGetTransferById(writer http.ResponseWriter, request *http.Request) {
	_, _ = writer.Write([]byte("{}"))
	log.Fatalln("[HandleGetTransferById] not implemented")
}