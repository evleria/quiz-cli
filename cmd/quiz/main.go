package main

import (
	"github.com/evleria/quiz-cli/pkg/cmd/root"
	"log"
)

func main() {
	cmd := root.NewRootCmd()
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
