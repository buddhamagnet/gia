package main

import (
	"log"
	"os"

	_ "github.com/buddhamagnet/gia/seeker/matchers"
	"github.com/buddhamagnet/gia/seeker/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	var term = "president"

	if len(os.Args) == 2 {
		term = os.Args[1]
	}

	search.Run(term)
}
