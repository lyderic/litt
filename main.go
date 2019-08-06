package main

import (
	"log"

	"github.com/lyderic/litt/cmd"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	cmd.Execute()
}
