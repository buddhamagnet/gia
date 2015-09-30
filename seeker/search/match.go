package search

type Result struct {
	Field   string
	Content string
}

type Matcher interface {
	Search(feed *Feed, term string) ([]*Result, error)
}
