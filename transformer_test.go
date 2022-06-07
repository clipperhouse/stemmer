package stemmer_test

import (
	"testing"

	"github.com/clipperhouse/stemmer"
	"golang.org/x/text/transform"
)

func TestStem(t *testing.T) {
	words := []string{
		"he",
		"intends",
		"to",
		"be",
		"doing",
		"some",
		"skiing",
	}

	for _, word := range words {
		stemmed, _, _ := transform.Bytes(stemmer.English, []byte(word))
		t.Logf("%q", stemmed)
	}
}
