package goheroku

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// structure of POST json
type helloworldPostRequest struct {
	Name string
}

type dataPostRequest struct {
	Field string
	Value string
}

// helloworldGetHandler handles GET method
func helloworldGetHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, HelloWorld())
}

// helloworldPostHandler handles POST method
func helloworldPostHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var hpr helloworldPostRequest
	err := decoder.Decode(&hpr)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, Hello(hpr.Name))
}

// dataGetHandler handles GET method
func dataGetHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	output := HelloWorld()

	field := r.URL.Query().Get("field")
	if field == "" {
		io.WriteString(w, output)
		return
	}

	value, err := GetData(field)
	if err != nil {
		log.Print(err)
		io.WriteString(w, output)
		return
	}

	output = value
	io.WriteString(w, output)
}

// dataPostHandler handles POST method
func dataPostHandler(w http.ResponseWriter, r *http.Request) {
	var token string
	tokens, ok := r.Header["Authorization"]
	if ok && len(tokens) > 0 {
		token = strings.TrimPrefix(tokens[0], "Bearer ")
	}

	if len(token) == 0 {
		log.Println("No token found")

		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "No token found\n")

		return
	}

	validToken, errT := CheckToken(token)
	if !validToken || errT != nil {
		log.Println(errT)

		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, errT.Error())

		return
	}

	decoder := json.NewDecoder(r.Body)

	var dpr dataPostRequest
	errD := decoder.Decode(&dpr)
	if errD != nil {
		log.Println(errD)

		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "Internal Server Error\n")

		return
	}

	done, errI := SetData(dpr.Field, dpr.Value)
	if !done || errI != nil {
		log.Println(errI)

		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, "Internal Server Error\n")

		return
	}

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "HTTP OK\n")
}

// statusNotFoundHandler handles unknown request method
func statusNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	io.WriteString(w, "Your requested method ("+r.Method+") is not found")
}

// helloworldRestAPIHandler is handles REST API (now only GET and POST) requests/responses
func helloworldRestAPIHandler(w http.ResponseWriter, r *http.Request) {
	// Set HTTP Response "Content-Type" header as JSON
	w.Header().Set("Content-type", "text/plain; charset=utf-8")

	// Set HTTP Response Body, based on HTTP request method
	switch r.Method {
	case "GET":
		helloworldGetHandler(w, r)
	case "POST":
		helloworldPostHandler(w, r)
	default:
		statusNotFoundHandler(w, r)
	}
}

// dataRestAPIHandler handles request response
func dataRestAPIHandler(w http.ResponseWriter, r *http.Request) {
	// Set HTTP Response "Content-Type" header as JSON
	w.Header().Set("Content-type", "text/plain; charset=utf-8")

	// Set HTTP Response Body, based on HTTP request method
	switch r.Method {
	case "GET":
		dataGetHandler(w, r)
	case "POST":
		dataPostHandler(w, r)
	default:
		statusNotFoundHandler(w, r)
	}
}

// HelloworldRestAPIWebserver provides connectivity for handling REST API request
func HelloworldRestAPIWebserver(pNum uint) {
	http.HandleFunc("/api/v1/helloworld", helloworldRestAPIHandler)
	http.HandleFunc("/api/v1/data", dataRestAPIHandler)

	log.Fatal(http.ListenAndServe(":"+fmt.Sprint(pNum), nil))
}
