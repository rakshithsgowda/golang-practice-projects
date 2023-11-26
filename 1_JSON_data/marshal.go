package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name     string `json:"firstname"`
	Age      int64
	Location string `json:"location,omitempty"`
}

func main() {
	

	person:=Person{
		Name:"indiana",
		Age:30,
	}

	personArray,err:= json.Marshal(person)
	if err!=nil{
		log.Fatal("Unable to Encode")
	}

	fmt.Println(string(personArray))
}