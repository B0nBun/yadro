package main

import (
	"log"
	"yadro-dns/server/cmd"
)

// TODO: Actual task with DNS and stuff
// TODO: Check out if HTTPS is needed

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
