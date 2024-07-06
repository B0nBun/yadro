package main

import (
	"log"
	"yadro/client/cmd"
)

// TODO: I don't really like the package structure and naming, so think about it later

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
