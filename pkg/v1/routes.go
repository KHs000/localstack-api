package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/KHs000/localstack-api/pkg/localsqs"
)

// SQSPong TODO
func SQSPong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, localsqs.Pong())
}

// CreateQueue TODO
func CreateQueue(w http.ResponseWriter, r *http.Request) {
	if !POST(r) {
		http.NotFound(w, r)
		return
	}

	dec := json.NewDecoder(r.Body)
	body := struct {
		QueueName string `json:"queueName"`
	}{}

	if err := dec.Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "could not parse json body")
		return
	}

	url, err := localsqs.Create(body.QueueName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, url)
}

// ListQueues TODO
func ListQueues(w http.ResponseWriter, r *http.Request) {
	if !GET(r) {
		http.NotFound(w, r)
		return
	}

	lst, err := localsqs.List()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "could not get queues list")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%v", lst))
}
