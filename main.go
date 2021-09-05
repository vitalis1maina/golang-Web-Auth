package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type person struct {
	First string
}

func main() {

// person1 := person{
// 	First: "vitalis",
// }
// person2 := person{
// 	First: "maina",
// }
// slicePerson := []person{person1, person2}
// bs, err := json.Marshal(slicePerson)
// if err != nil {
// 	log.Panic(err) //panic when u get a programmer error
// }
// fmt.Println("PRINT JSON:", string(bs))
// slicePerson2 := []person{}
// err = json.Unmarshal(bs, &slicePerson2)
// if err != nil {
// 	log.Panic(err)
// }
// fmt.Println("Unmashalled data:", slicePerson2)

	htttp.Handlefunc("/encode",encode)
	http.HandleFunc("/decode",decode)
	http.ListenAndServe(":8080",nil)
}
func encode(http.ResponseWriter r *http.Request){

}
func decode(http.ResponseWriter r * http.Request){
	
}
