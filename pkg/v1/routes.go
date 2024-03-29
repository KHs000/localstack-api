package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/KHs000/localstack-api/pkg/localkinesis"
	"github.com/KHs000/localstack-api/pkg/localsqs"
)

var (
	sqsClient localsqs.Client
	knsClient localkinesis.Client
)

func init() {
	sqsClient = localsqs.NewClient()
	knsClient = localkinesis.NewClient()
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

	url, err := sqsClient.Create(body.QueueName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, url)
}

// GetQueueAttributes TODO
func GetQueueAttributes(w http.ResponseWriter, r *http.Request) {
	if !POST(r) {
		http.NotFound(w, r)
		return
	}

	dec := json.NewDecoder(r.Body)
	body := struct {
		QueueURL   string   `json:"queueUrl" required:"true"`
		Attributes []string `json:"attributes"`
	}{}

	if err := dec.Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "could not parse json body")
		return
	}

	if body.Attributes == nil || len(body.Attributes) == 0 {
		body.Attributes = []string{"All"}
	}

	attr, err := sqsClient.GetAttributes(body.QueueURL, body.Attributes...)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "could not get queue attributes")
		return
	}
	for k, v := range attr {
		fmt.Fprintf(w, "%s: %s\n", k, v)
	}
}

// ListQueues TODO
func ListQueues(w http.ResponseWriter, r *http.Request) {
	if !GET(r) {
		http.NotFound(w, r)
		return
	}

	lst, err := sqsClient.List()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "could not get queues list")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%v", lst))
}

// PurgeQueue TODO
func PurgeQueue(w http.ResponseWriter, r *http.Request) {
	if !POST(r) {
		http.NotFound(w, r)
		return
	}

	dec := json.NewDecoder(r.Body)
	body := struct {
		QueueURL string `json:"queueUrl"`
	}{}

	if err := dec.Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "could not parse json body")
		return
	}

	if err := sqsClient.Purge(body.QueueURL); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, "queue purged")
}

// CreateStream TODO
func CreateStream(w http.ResponseWriter, r *http.Request) {
	if !POST(r) {
		http.NotFound(w, r)
		return
	}

	dec := json.NewDecoder(r.Body)
	body := struct {
		StreamName string `json:"streamName"`
	}{}

	if err := dec.Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "could not parse json body")
		return
	}

	if err := knsClient.Create(body.StreamName); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "stream created")
}

// ListStreams TODO
func ListStreams(w http.ResponseWriter, r *http.Request) {
	if !GET(r) {
		http.NotFound(w, r)
		return
	}

	nms, err := knsClient.List()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	for _, v := range nms {
		fmt.Fprintf(w, "%s\n", v)
	}
}

// PutRecord TODO
func PutRecord(w http.ResponseWriter, r *http.Request) {
	if !POST(r) {
		http.NotFound(w, r)
		return
	}

	dec := json.NewDecoder(r.Body)
	body := struct {
		Data       string `json:"data"`
		StreamName string `json:"streamName"`
	}{}

	if err := dec.Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "could not parse json body")
		return
	}

	if err := knsClient.PutRecord([]byte(body.Data), body.StreamName); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "record sent")
}
