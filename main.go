package main

import (
	"github.com/felipeas314/go/database"
	"github.com/felipeas314/go/routers"
	"log"
)

func main() {
	database.Setup()
	r := routers.Setup()
	if err := r.Run("127.0.0.1:8081"); err != nil {
		log.Fatal(err)
	}
}