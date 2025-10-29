package parser

import (
	"archive/zip"
	"bytes"
	"io"
	"io/ioutil"
)

type docxParser struct{}

func (d *docxParser) SupportedTypes() []string {
	return []string{"application/vnd.openxmlformats-officedocument.wordprocessingml.document"}
}

func (d *docxParser) ExtractText(r io.Reader) (string, error) {
	var buf bytes.Buffer
	io.Copy(&buf, r)

	zr, err := zip.NewReader(bytes.NewReader(buf.Bytes()), int64(buf.Len()))
	if err != nil {
		return "", err
	}

	for _, f := range zr.File {
		if f.Name == "word/document.xml" {
			rc, err := f.Open()
			if err != nil {
				return "", err
			}
			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return "", err
			}
			return string(data), nil
		}
	}

	return "", nil
}

func (d *docxParser) ExtractMetadata(r io.Reader) (Metadata, error) {
	return Metadata{"MimeType": d.SupportedTypes()[0]}, nil
}

func init() {
	Register(&docxParser{})
}
