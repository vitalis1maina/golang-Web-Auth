package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	Name string
}

func main() {

	http.HandleFunc("/encode", encode)
	http.HandleFunc("/decode", decode)
	http.ListenAndServe(":8080", nil)

}
func encode(w http.ResponseWriter, r *http.Request) {
	p1 := person{
		"vitalis",
	}
	p2 := person{
		"maina",
	}
	people := []person{p1, p2}

	err := json.NewEncoder(w).Encode(people)
	if err != nil {
		log.Println("bad data to encode", err)
	}

}
func decode(w http.ResponseWriter, r *http.Request) {
	people := []person{}
	err := json.NewDecoder(r.Body).Decode(&people)
	if err != nil {
		log.Println("decoding bad data:", err)
	}
	log.Println(people)
}
