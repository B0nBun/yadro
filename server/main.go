package main

import (
	"log"
	"yadro-dns/server/cmd"
)

// TODO: Actual task with DNS and stuff

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
