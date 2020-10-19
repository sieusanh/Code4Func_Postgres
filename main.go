package main

import (
	"fmt"
	"gmodule/driver"
)

// host, port, user, password, dbname
const (
	host = "localhost"
	port = "5432"
	user = "postgres"
	password = "123456"
	dbname = "code4func"
)

func main() {
	db := driver.Connect(host, port, user, password, dbname)

	err := db.SQL.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("connection ok")
}