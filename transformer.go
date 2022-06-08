package stemmer

import (
	"sync"

	"github.com/blevesearch/snowballstem"
	"golang.org/x/text/transform"
)

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

// looks like a premature optimization, maybe it is, but env
// is stateful and I don't want Transformer to be
var envPool = sync.Pool{
	New: func() interface{} {
		return snowballstem.NewEnv("")
	},
}

func (t *Transformer) Transform(dst, src []byte, atEOF bool) (nDst, nSrc int, err error) {
	env := envPool.Get().(*snowballstem.Env)

	env.SetCurrent(string(src))
	nSrc = len(src)

	t.stem(env)

	stemmed := []byte(env.Current())
	nDst = len(stemmed)

	copy(dst, src)

	envPool.Put(env)
	return nDst, nSrc, err
}

func (t *Transformer) Reset() {}
