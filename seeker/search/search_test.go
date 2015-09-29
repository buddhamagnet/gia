package search_test

import (
	"testing"

	"github.com/buddhamagnet/gia/seeker/search"
)

func TestSearch(t *testing.T) {
	result := search.Run("testing")
	if result != "term: testing" {
		t.Errorf("expected to see 'term: testing', got %s\n", result)
	}
}
