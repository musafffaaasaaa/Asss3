package main

import (
	"ass3adv2/pkg/store/postgres"
	"fmt"
	"log"
)

func main() {
	_, err := postgres.OpenDb("localhost", "8520", "postgres", "123456", "ass3adv2")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Success connect")
}
