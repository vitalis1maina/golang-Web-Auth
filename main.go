package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First string
}

func main() {

	person1 := person{
		First: "vitalis",
	}
	person2 := person{
		First: "maina",
	}
	slicePerson := []person{person1, person2}
	bs, err := json.Marshal(slicePerson)
	if err != nil {
		log.Panic(err) //panic when u get a programmer error
	}
	fmt.Println(string(bs))
}