package search

import (
	"fmt"
	"log"
)

type Result struct {
	Field   string
	Content string
}

func (r Result) String() string {
	return fmt.Sprintf("%s:\n%s\n\n", r.Field, r.Content)
}

type Matcher interface {
	Search(feed *Feed, term string) ([]*Result, error)
}

func Match(matcher Matcher, feed *Feed, term string, results chan<- *Result) {
	resultSet, err := matcher.Search(feed, term)
	if err != nil {
		log.Println(err)
		return
	}
	for _, result := range resultSet {
		results <- result
	}
	if len(results) == 0 {
		fmt.Printf("no results found at %s\n", feed.Name)
	}
}

func Display(results chan *Result) {
	for result := range results {
		fmt.Printf("%s", result)
	}
}
