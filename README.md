[![Documentation](https://pkg.go.dev/badge/github.com/clipperhouse/stemmer.svg)](https://pkg.go.dev/github.com/clipperhouse/stemmer)

Package `stemmer` is a set of Snowball language stemmers, implemented as `transform.Transformer` (see [golang.org/x/text/](https://golang.org/x/text/)). Implementing `transform.Transformer` allows use with the `x/text` ecosystem.

My narrower motivation is to offer stemming in [this Unicode text segmenter](https://pkg.go.dev/github.com/clipperhouse/uax29/words). To use a stemmer in that package, you will call the `Transform` method on a `Segmenter` or `Scanner`, along the lines of:

```go
import "github.com/clipperhouse/stemmer"

text := []byte("I've got my skis and I am heading to the mountains.")
segmenter := words.NewSegmenter(text)
segmenter.Transform(stemmer.English)

// do stuff with segmenter
```

This `stemmer` package is a thin layer over [this Snowball package](https://github.com/blevesearch/snowballstem), thanks to Bleve.
