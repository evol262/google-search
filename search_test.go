package googlesearch_test

import (
	"context"
	"testing"

	"github.com/rocketlaunchr/google-search"
)

var ctx = context.Background()

func TestSearch(t *testing.T) {

	q := "Hello World"

	opts := googlesearch.SearchOptions{
		Limit: 20,
	}

	returnLinks, err := googlesearch.Search(ctx, q, opts)
	if err != nil {
		t.Errorf("something went wrong: %v", err)
		return
	}

	if len(returnLinks) == 0 {
		t.Errorf("no results returned: %v", returnLinks)
	}
}

func TestNoResults(t *testing.T) {
	q := "\"Alexandra's Purse Fund\"+\"LGBT\""

	opts := googlesearch.SearchOptions{
		Limit: 1,
	}

	returnLinks, err := googlesearch.Search(ctx, q, opts)
	if err != nil {
		return
	}

	if len(returnLinks) >= 0 {
		t.Errorf("Results returned: %v", returnLinks)
	}
}
