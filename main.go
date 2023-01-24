package main

import (
	"log"

	"github.com/squarefactory/submer-pod-exporter/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalln("couldn't execute command", err)
	}
}
