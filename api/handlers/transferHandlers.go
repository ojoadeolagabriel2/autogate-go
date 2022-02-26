package handlers

import (
	qrcode "github.com/skip2/go-qrcode"
	"log"
	"net/http"
)

func GetTransferById(writer http.ResponseWriter, request *http.Request) {
	_, _ = writer.Write([]byte("{}"))
	log.Fatalln("[GetTransferById] not implemented")
}

func GetQrCode(writer http.ResponseWriter, request *http.Request) {
	if png, err := qrcode.Encode("https://localhost/v1/transactions/1", qrcode.Medium, 256); err == nil {
		writer.Header().Set("Content-Disposition", "attachment; filename=foo.pdf")
		_, err := writer.Write(png)
		if err != nil {
			log.Fatal("could not complete [GetQrCode]")
		}
	}
}