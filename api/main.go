package main

import (
	"api/db"
	"fmt"
)

func main() {
	db, err := db.NewDB()
	if err != nil {
		panic(err)
	}

	fmt.Println("DB connection successfully opened")
	fmt.Println(db)
}
