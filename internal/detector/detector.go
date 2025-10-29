package detector

import (
	"io"
	"net/http"
)

type Detector interface {
	Detect(r io.Reader) (mimeType string, err error)
}

type httpDetector struct{}

func (d *httpDetector) Detect(r io.Reader) (string, error) {
	buf := make([]byte, 512)
	n, _ := r.Read(buf)
	return http.DetectContentType(buf[:n]), nil
}

// Default detector
var Default Detector = &httpDetector{}
