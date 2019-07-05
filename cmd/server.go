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
	knsHandler struct {
		BaseRoute string
	}
)

const (
	v1Prefix  = "/v1"
	sqsPrefix = "/sqs"
	knsPrefix = "/kns"
)

var (
	v1sqsPrefix = v1Prefix + sqsPrefix
	v1knsPrefix = v1Prefix + knsPrefix
)

func (h sqsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != h.BaseRoute {
		http.NotFound(w, r)
		return
	}
}

func (k knsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != k.BaseRoute {
		http.NotFound(w, r)
		return
	}
}

func main() {
	v1mux := http.NewServeMux()

	// SQS routes
	v1mux.Handle(v1sqsPrefix, sqsHandler{BaseRoute: v1sqsPrefix})
	v1mux.HandleFunc(v1sqsPrefix+"/create", v1.CreateQueue)
	v1mux.HandleFunc(v1sqsPrefix+"/attributes", v1.GetQueueAttributes)
	v1mux.HandleFunc(v1sqsPrefix+"/list", v1.ListQueues)
	v1mux.HandleFunc(v1sqsPrefix+"/purge", v1.PurgeQueue)

	// Kinesis routes
	v1mux.Handle(v1knsPrefix, knsHandler{BaseRoute: v1knsPrefix})
	v1mux.HandleFunc(v1knsPrefix+"/create", v1.CreateStream)
	v1mux.HandleFunc(v1knsPrefix+"/list", v1.ListQueues)

	log.Println("Server up...")
	http.ListenAndServe(":8082", v1mux)
}
