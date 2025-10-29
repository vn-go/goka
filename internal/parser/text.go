package parser

import (
	"io"
	"io/ioutil"
)

type textParser struct{}

func (t *textParser) SupportedTypes() []string {
	return []string{"text/plain"}
}

func (t *textParser) ExtractText(r io.Reader) (string, error) {
	data, err := ioutil.ReadAll(r)
	return string(data), err
}

func (t *textParser) ExtractMetadata(r io.Reader) (Metadata, error) {
	return Metadata{"MimeType": "text/plain"}, nil
}

func init() {
	Register(&textParser{})
}
