package main

import (
	"flag"
	"log"
	"os"

	_ "github.com/buddhamagnet/gia/seeker/matchers"
	"github.com/buddhamagnet/gia/seeker/search"
)

var term string

func init() {
	log.SetOutput(os.Stdout)
	flag.StringVar(&term, "term", "golang", "the term to search for")
}

func main() {
	flag.Parse()

	if len(os.Args) == 2 {
		term = os.Args[1]
	}

	search.Run(term)
}
