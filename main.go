package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// Message is used for storing the
type Message struct {
	Hostname    string    `json:"hostname"`
	ServiceName string    `json:"service_name"`
	Error       string    `json:"error"`
	Messages    []Message `json:"messages"`
}

var hostname string

func splitPath(path string) (string, string) {
	if paths := strings.SplitN(path, "/", 2); len(paths) > 1 {
		fmt.Println("Splitting", paths)
		return paths[0], paths[1]
	}
	return path, ""
}

// NewMessage creates a new Message
func NewMessage(path string) Message {
	msg := Message{
		Hostname:    hostname,
		ServiceName: os.Getenv("SERVICE_NAME"),
		Messages:    make([]Message, 0),
	}

	if service, path := splitPath(path); service != "" {
		fmt.Println("Service:Path", service, path)
		msg.Messages = append(msg.Messages, fetchSubURL(service, path))
	}

	return msg
}

func fetchSubURL(service, path string) Message {
	url := fmt.Sprintf("http://%s.service.consul:3000/%s", service, path)
	fmt.Println("Querying:", url)
	msg := Message{}

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Failed request:", err)
		msg.Error = err.Error()
		return msg
	}

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&msg); err != nil {
		fmt.Println("Failed Decoding Response:", err)
	}

	return msg
}

func handleRequest(w http.ResponseWriter, req *http.Request) {
	m := NewMessage(req.URL.Path[1:])
	json := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	json.Encode(m)
}

func handleIcon(w http.ResponseWriter, req *http.Request) {
}

func main() {
	var err error
	if hostname, err = os.Hostname(); err != nil {
		fmt.Println("Unable to find Hostname")
		os.Exit(1)
	}

	http.HandleFunc("/", handleRequest)
	http.HandleFunc("/favicon.ico", handleIcon)
	http.ListenAndServe(":3000", nil)
}
