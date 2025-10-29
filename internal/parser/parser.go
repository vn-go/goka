package parser

import "io"

type Parser interface {
	SupportedTypes() []string
	ExtractText(r io.Reader) (string, error)
	ExtractMetadata(r io.Reader) (Metadata, error)
}
