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
	search.Run("president")
}
