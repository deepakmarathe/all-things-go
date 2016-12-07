package main

import (
	//"encoding/json"
	"log"
	"net/http"
	"io/ioutil"
	"fmt"
	//"encoding/json"
)

type test_struct struct {
	Test string
}

func test(rw http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("---------------error---------")
	}
	log.Println(string(body))
	//var t test_struct
	//err = json.Unmarshal(body, &t)
	//if err != nil {
	//	fmt.Println("---------------error---------")
	//}
	//log.Println(t.Test)
}

func main() {
	http.HandleFunc("/bookinglog", test)
	log.Fatal(http.ListenAndServe(":8080", nil))
}