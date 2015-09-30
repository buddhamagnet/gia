package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

func Match(matcher Matcher, feed *Feed, term string, results chan *Result) {}

func Display(chan *Result) {}

func Run(term string) {
	// Get the feeds.
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}
	// We need a waitgroup and a channel of results.
	var wg sync.WaitGroup
	results := make(chan *Result)
	// Add the length of the feeds to the waitgroup.
	wg.Add(len(feeds))
	// Iterate through the feeds and select the matcher.
	for _, feed := range feeds {
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}
		// Fire up a goroutine for the search and
		// decrement the waitgroup when done.
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, term, results)
			wg.Done()
		}(matcher, feed)
	}

	// Fire up a goroutine to monitor when all
	// work has been completed.
	go func() {
		wg.Wait()
		close(results)
	}()
	// Start to display results as they come in.
	Display(results)
}
