package search

type defaultMatcher struct{}

func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

func (d defaultMatcher) Search(feed *Feed, term string) ([]*Result, error) {
	return nil, nil
}

func Register(label string, matcher Matcher) {
	matchers[label] = matcher
}
