package main

import (
	"log"
	"yadro-dns/client/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
