package main

import (
	"fmt"
	"log"

	"github.com/vn-go/goka"
)

func main() {
	text, err := goka.ExtractTextFromFile("sample.pdf")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Extracted text:")
	fmt.Println(text)

	meta, err := goka.ExtractMetadataFromFile("sample.pdf")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Metadata:", meta)
}
