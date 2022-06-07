package stemmer_test

import (
	"testing"

	"github.com/clipperhouse/stemmer"
	"golang.org/x/text/transform"
)

func TestStem(t *testing.T) {
	input := []string{
		"He",
		"intends",
		"to",
		"be",
		"doing",
		"some",
		"skiing",
		"and",
		"imbibing",
		"Açaí",
		"deliciously",
		".",
	}

	expected := []string{
		"He",
		"intend",
		"to",
		"be",
		"do",
		"some",
		"ski",
		"and",
		"imbib",
		"Açaí",
		"delici",
		".",
	}

	for i, word := range input {
		stemmed, _, err := transform.Bytes(stemmer.English, []byte(word))
		if err != nil {
			t.Fatal(err)
		}

		if string(stemmed) != expected[i] {
			t.Fatalf("expected %q, got %q", expected[i], stemmed)

		}
	}
}
