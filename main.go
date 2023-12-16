package main

import (
	"fmt"
	"log"
)

func main() {
	path := "kodi.db"

	db, err := Open(path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("db: %+v", db)
}
