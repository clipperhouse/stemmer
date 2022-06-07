package stemmer

import (
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

func (t *Transformer) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	env := snowballstem.NewEnv(string(src))

	nSrc = len(src)

	t.stem(env)

	stemmed := env.Current()
	nDst = len(stemmed)

	for i := range stemmed {
		dst[i] = stemmed[i]
	}

	return nDst, nSrc, nil
}

func (t *Transformer) Reset() {}
