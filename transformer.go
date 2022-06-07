package stemmer

import (
	"errors"
	"fmt"
	"unicode/utf8"

	"github.com/blevesearch/snowballstem"
	"golang.org/x/text/transform"
)

// ensure Transformer implements transform.Transformer from golang.org/x/text/
//lint:ignore U1000 unused
var check transform.Transformer = &Transformer{}

type StemFunc func(*snowballstem.Env) bool

func NewTransformer(stem StemFunc) *Transformer {
	return &Transformer{
		stem: stem,
	}
}

type Transformer struct {
	stem StemFunc
}

var ErrInvalidUtf8 = errors.New("stemming resulted in invalid UTF-8")

func (t *Transformer) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	env := snowballstem.NewEnv(string(src))

	nSrc = len(src)

	t.stem(env)

	stemmed := []byte(env.Current())
	nDst = len(stemmed)

	if !utf8.Valid(stemmed) {
		err = fmt.Errorf("stemmed %q, resulting in %q: %w", string(src), stemmed, ErrInvalidUtf8)
	}

	copy(dst, src)

	return nDst, nSrc, err
}

func (t *Transformer) Reset() {}
