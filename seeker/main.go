package main

import (
	"log"
	"os"

	"github.com/buddhamagnet/gia/seeker/search"
)

func init() {
	log.SetOutput(os.Stdout)
}
func main() {
	search.Run("president")
}
