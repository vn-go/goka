package goka

import (
	"fmt"
	"os"

	"github.com/vn-go/goka/internal/detector"
	"github.com/vn-go/goka/internal/parser"
)

func ExtractTextFromFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	mime, err := detector.Default.Detect(f)
	if err != nil {
		return "", err
	}

	p := parser.GetByMime(mime)
	if p == nil {
		return "", fmt.Errorf("no parser for %s", mime)
	}

	f.Seek(0, 0)
	return p.ExtractText(f)
}

func ExtractMetadataFromFile(path string) (parser.Metadata, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	mime, err := detector.Default.Detect(f)
	if err != nil {
		return nil, err
	}

	p := parser.GetByMime(mime)
	if p == nil {
		return nil, fmt.Errorf("no parser for %s", mime)
	}

	f.Seek(0, 0)
	return p.ExtractMetadata(f)
}
