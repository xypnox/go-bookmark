package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	GetLinksSearch("dhc")
	port := ":8080"
	mux := http.DefaultServeMux

	mux.HandleFunc("/create", CreateHandler)
	mux.HandleFunc("/fetch/link", FetchLinkHandler)
	mux.HandleFunc("/fetch/tag", FetchTagHandler)

	fmt.Println("Server Is Starting")
	err := http.ListenAndServe(port, mux)
	if err != nil {
		fmt.Println("Your server is down")
	}
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	var rqst CreateRequest

	decoder := json.NewDecoder(r.Body)
	err1 := decoder.Decode(&rqst)

	if err1 != nil {
		log.Fatal(err1)
	}

	fmt.Println(rqst.URL)
	fmt.Println(rqst.Tag)
	fmt.Println(rqst.CreateTime)

	link := LinkType{
		URL:        rqst.URL,
		Title:      rqst.Title,
		Tag:        rqst.Tag,
		CreateTime: rqst.CreateTime,
	}

	var status bool

	// Save the Request if it doesn't exists else return false status

	StoreLink(link)

	status = true

	response := CreateResponse{
		Message: status,
	}
	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func FetchLinkHandler(w http.ResponseWriter, r *http.Request) {
	var rqst FetchLinkRequest

	decoder := json.NewDecoder(r.Body)
	err1 := decoder.Decode(&rqst)

	if err1 != nil {
		log.Fatal(err1)
	}

	var status bool

	// Find the link in the database

	links := GetLinksSearch(rqst.SearchTerm)

	status = true

	response := FetchLinkResponse{
		Message: status,
		Links:   links,
	}

	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func FetchTagHandler(w http.ResponseWriter, r *http.Request) {
	var rqst FetchTagRequest

	decoder := json.NewDecoder(r.Body)
	err1 := decoder.Decode(&rqst)
	if err1 != nil {
		log.Fatal(err1)
	}

	var status bool
	var size = 0

	status = true
	links := GetLinksTag(rqst.Tag)
	size = len(links)

	if size == 0 {
		status = false
	}

	response := FetchTagResponse{
		Message: status,
		Size:    size,
		Links:   links,
	}

	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
