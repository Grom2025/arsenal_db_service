package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// handle route using handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			// in case of any error
			log.Fatalln("Parse Error ", r)
		}
		isGroup := r.Form.Get("PAR")
		name := r.Form.Get("Name")
		code := r.Form.Get("Code")
		group := r.Form.Get("Group")
		var desc, price, qantity string

		if isGroup == "0" {
			desc = r.Form.Get("Desc")
			price = r.Form.Get("Price")
			qantity = r.Form.Get("Qantity")
		}
		// x will be "" if parameter is not set
		fmt.Println(code, group, name, desc, price, qantity)
		//fmt.Fprintf(w, "Welcome to new server!")
	})

	// listen to port
	http.ListenAndServe(":8000", nil)

}
