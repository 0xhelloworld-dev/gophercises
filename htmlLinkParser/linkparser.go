package linkparser

import (
	"io"

	html "golang.org/x/net/html"
)

func TokenizeFile(filename *io.Reader) {
	html.NewTokenizer(filename)
}
