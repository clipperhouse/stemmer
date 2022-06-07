// Package stemmer set of Snowball language stemmers, implemented as transform.Transformer (from golang.org/x/text/).
package stemmer

import (
	"github.com/blevesearch/snowballstem/arabic"
	"github.com/blevesearch/snowballstem/danish"
	"github.com/blevesearch/snowballstem/dutch"
	"github.com/blevesearch/snowballstem/english"
	"github.com/blevesearch/snowballstem/finnish"
	"github.com/blevesearch/snowballstem/french"
	"github.com/blevesearch/snowballstem/german"
	"github.com/blevesearch/snowballstem/hungarian"
	"github.com/blevesearch/snowballstem/irish"
	"github.com/blevesearch/snowballstem/italian"
	"github.com/blevesearch/snowballstem/norwegian"
	"github.com/blevesearch/snowballstem/porter"
	"github.com/blevesearch/snowballstem/portuguese"
	"github.com/blevesearch/snowballstem/romanian"
	"github.com/blevesearch/snowballstem/russian"
	"github.com/blevesearch/snowballstem/spanish"
	"github.com/blevesearch/snowballstem/swedish"
	"github.com/blevesearch/snowballstem/tamil"
	"github.com/blevesearch/snowballstem/turkish"
	"golang.org/x/text/transform"
)

// Arabic is a Snowball stemmer for Arabic
var Arabic transform.Transformer = NewTransformer(arabic.Stem)

// Danish is a Snowball stemmer for Danish
var Danish transform.Transformer = NewTransformer(danish.Stem)

// Dutch is a Snowball stemmer for Dutch
var Dutch transform.Transformer = NewTransformer(dutch.Stem)

// English is a Snowball stemmer for English
var English transform.Transformer = NewTransformer(english.Stem)

// Finnish is a Snowball stemmer for Finnish
var Finnish transform.Transformer = NewTransformer(finnish.Stem)

// French is a Snowball stemmer for French
var French transform.Transformer = NewTransformer(french.Stem)

// German is a Snowball stemmer for German
var German transform.Transformer = NewTransformer(german.Stem)

// Hungarian is a Snowball stemmer for Hungarian
var Hungarian transform.Transformer = NewTransformer(hungarian.Stem)

// Irish is a Snowball stemmer for Irish
var Irish transform.Transformer = NewTransformer(irish.Stem)

// Italian is a Snowball stemmer for Italian
var Italian transform.Transformer = NewTransformer(italian.Stem)

// Norwegian is a Snowball stemmer for Norwegian
var Norwegian transform.Transformer = NewTransformer(norwegian.Stem)

// Porter is a Porter stemmer
var Porter transform.Transformer = NewTransformer(porter.Stem)

// Portuguese is a Snowball stemmer for Portuguese
var Portuguese transform.Transformer = NewTransformer(portuguese.Stem)

// Romanian is a Snowball stemmer for Romanian
var Romanian transform.Transformer = NewTransformer(romanian.Stem)

// Russian is a Snowball stemmer for Russian
var Russian transform.Transformer = NewTransformer(russian.Stem)

// Spanish is a Snowball stemmer for Spanish
var Spanish transform.Transformer = NewTransformer(spanish.Stem)

// Swedish is a Snowball stemmer for Swedish
var Swedish transform.Transformer = NewTransformer(swedish.Stem)

// Tamil is a Snowball stemmer for Tamil
var Tamil transform.Transformer = NewTransformer(tamil.Stem)

// Turkish is a Snowball stemmer for Turkish
var Turkish transform.Transformer = NewTransformer(turkish.Stem)
