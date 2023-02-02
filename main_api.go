package main

import (
	"fmt"
	"log"
	http "net/http"
)

func readDataFrom1C(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		// in case of any error
		log.Fatalln("Parse Error ", r)
	}
	isGroup := r.Form.Get("PAR")
	name := r.Form.Get("Name")
	code := r.Form.Get("Code")
	group := r.Form.Get("Group")

	desc := r.Form.Get("Desc")
	price := r.Form.Get("Price")
	quantity := r.Form.Get("Quantity")

	// debug code
	log.Println(isGroup, code, group, name, desc, price, quantity)

	_, err = fmt.Fprintf(w, "OK")
	if err != nil {
		return
	}

}

func main() {
	// handle route using handler function
	http.HandleFunc("/", readDataFrom1C)

	// listen to port
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Println("Server launch error")
	}

}
