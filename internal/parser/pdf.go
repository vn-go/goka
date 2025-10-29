package parser

import (
	"bytes"
	"io"

	"github.com/ledongthuc/pdf"
)

type pdfParser struct{}

func (p *pdfParser) SupportedTypes() []string {
	return []string{"application/pdf"}
}

func (p *pdfParser) ExtractText(r io.Reader) (string, error) {
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, r); err != nil {
		return "", err
	}

	content, err := pdf.NewReader(bytes.NewReader(buf.Bytes()), int64(buf.Len()))
	if err != nil {
		return "", err
	}

	textIO, err := content.GetPlainText()
	if err != nil {
		return "", err
	}

	data, err := io.ReadAll(textIO)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (p *pdfParser) ExtractMetadata(r io.Reader) (Metadata, error) {
	return Metadata{"MimeType": "application/pdf"}, nil
}

func init() {
	Register(&pdfParser{})
}
