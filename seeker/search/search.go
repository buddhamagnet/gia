package search

import (
	"fmt"
	_ "log"
	_ "sync"
)

type Matcher struct{}

var matchers = make(map[string]Matcher)

func Run(term string) string {
	return fmt.Sprintf("term: %s", term)
}
