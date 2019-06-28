package main

import (
	"log"
	"net/http"

	v1 "github.com/KHs000/localstack-api/pkg/v1"
)

type (
	sqsHandler struct {
		BaseRoute string
	}
)

const (
	v1Prefix  = "/v1"
	sqsPrefix = "/sqs"
)

var (
	v1sqsPrefix = v1Prefix + sqsPrefix
)

func (h sqsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != h.BaseRoute {
		http.NotFound(w, r)
		return
	}
}

func main() {
	v1mux := http.NewServeMux()

	v1mux.Handle(v1sqsPrefix, sqsHandler{BaseRoute: v1sqsPrefix})
	v1mux.HandleFunc(v1sqsPrefix+"/ping", v1.SQSPong)
	v1mux.HandleFunc(v1sqsPrefix+"/create", v1.CreateQueue)
	v1mux.HandleFunc(v1sqsPrefix+"/list", v1.ListQueues)

	log.Println("Server up...")
	http.ListenAndServe(":8082", v1mux)
}
