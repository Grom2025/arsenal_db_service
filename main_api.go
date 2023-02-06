package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"log"
	http "net/http"
	"os"
)

var conn *pgx.Conn

func readDataFrom1C(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		// in case of any error
		log.Fatalln("Parse Error ", r, "error: ", err)
		return
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

	sqlString := ""

	if isGroup == "0" {
		sqlString = fmt.Sprintf(
			"INSERT INTO ars.items (id , fname, fgroup, fdesc, price, quantity) "+
				"VALUES('%s','%s','%s','%s',%s,'%s') "+
				"ON CONFLICT (id) "+
				"DO "+
				"UPDATE SET fname = '%s', fgroup = '%s', fdesc = '%s' ,"+
				"price = %s, quantity = '%s' ",
			code, name, group, desc, price, quantity,
			name, group, desc, price, quantity)
	} else {

		sqlString = fmt.Sprintf(
			"INSERT INTO ars.tgroups (id , fname, fgroup) "+
				"VALUES('%s','%s','%s') "+
				"ON CONFLICT (id) "+
				"DO "+
				"UPDATE SET fname = '%s', fgroup = '%s' ",
			code, name, group,
			name, group)
	}

	log.Println(sqlString)

	_, err = conn.Exec(context.Background(), sqlString)
	if err != nil {
		log.Fatalln("SQL Error ", "error: ", err)
		return
	}

	_, err = fmt.Fprintf(w, "OK")
	if err != nil {
		return
	}

}

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		return
	}
	conn, err = pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Println(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer func(conn *pgx.Conn, ctx context.Context) {
		err = conn.Close(ctx)
		if err != nil {

		}
	}(conn, context.Background())

	// handle route using handler function
	http.HandleFunc("/", readDataFrom1C)

	// listen to port
	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Println("Server launch error")
	}

}
