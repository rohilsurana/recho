package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	handlerFunc := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		req, err := httputil.DumpRequest(request, true)
		if err != nil {
			log.Println("error dumping request:", err)
			writer.WriteHeader(http.StatusBadRequest)
		}
		_, err = writer.Write(req)
		if err != nil {
			log.Println("error writing to request: ", err)
			writer.WriteHeader(http.StatusInternalServerError)
		}
	})

	if err := http.ListenAndServe(":8080", handlerFunc); err != nil {
		log.Fatal("error server stopped: ", err)
	}
}
