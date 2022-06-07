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
)

var Arabic = NewTransformer(arabic.Stem)
var Danish = NewTransformer(danish.Stem)
var Dutch = NewTransformer(dutch.Stem)
var English = NewTransformer(english.Stem)
var Finnish = NewTransformer(finnish.Stem)
var French = NewTransformer(french.Stem)
var German = NewTransformer(german.Stem)
var Hungarian = NewTransformer(hungarian.Stem)
var Irish = NewTransformer(irish.Stem)
var Italian = NewTransformer(italian.Stem)
var Norwegian = NewTransformer(norwegian.Stem)
var Porter = NewTransformer(porter.Stem)
var Portuguese = NewTransformer(portuguese.Stem)
var Romanian = NewTransformer(romanian.Stem)
var Russian = NewTransformer(russian.Stem)
var Spanish = NewTransformer(spanish.Stem)
var Swedish = NewTransformer(swedish.Stem)
var Tamil = NewTransformer(tamil.Stem)
var Turkish = NewTransformer(turkish.Stem)
