package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"log"
)

type Animal struct {
	ID         int
	AnimalType string
	Nickname   string
	Zone       int
	Age        int
}


func main() {
	data :=&Animal{
		AnimalType: "Velociraptor",
		Nickname: "patro",
		Zone: 3,
		Age: 10,
	}

	var b bytes.Buffer
	json.NewEncoder(&b).Encode(data)


	//resp, err := http.Post("http://localhost:8080/api/dinos/add", "application/json", &b)
	//if err != nil || resp.StatusCode != 200 {
	//	log.Fatal(err)
	//}

	resp, err := http.Post("http://localhost:8080/api/dinos/edit/patro", "application/json", &b)
	if err != nil || resp.StatusCode != 200 {
		log.Fatal(resp.Status, err)
	}
}
