package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Tag object
type Tag struct {
	Goodreads_book_id string `json:"goodreads_book_id,omitempty"`
	Tag_id            string `json:"tag_id,omitempty"`
	Count             int    `json:"count,omitempty"`
}

// Tag structure (array of tags)
var Tags []Tag
var TagsResp []Tag

// GetTag function to serve a single tag
func GetTag(w http.ResponseWriter, r *http.Request) {
	resp := TagsResp
	params := mux.Vars(r)
	for _, item := range Tags {
		if item.Tag_id == params["id"] {
			resp = append(resp, item)
		}
	}
	json.NewEncoder(w).Encode(resp)
}

// GetAll function to serve all tags
func GetAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Tags)
}

// main function to boot up everything
func main() {
	//json loading
	jsonFile, err := ioutil.ReadFile("book_tags.json")
	if err != nil {
		fmt.Print(err)
	}
	err = json.Unmarshal(jsonFile, &Tags)
	if err != nil {
		fmt.Println("error:", err)
	}

	// api logic
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	router.HandleFunc("/tags/{id}", GetTag).Methods("GET")
	router.HandleFunc("/tags-all", GetAll).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
