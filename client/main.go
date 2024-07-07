package main

import (
	"os"

	"yadro/client/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		// Presumably the error was already printed by cobra, so we just exit
		os.Exit(1)
	}
}
