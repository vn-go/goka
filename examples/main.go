package main

import (
	"fmt"
	"log"

	"github.com/vn-go/goka"
)

func main() {
	test_file := `D:\code\go\vn-go-dx\goka\test-data\UQ 678.pdf`
	text, err := goka.ExtractTextFromFile(test_file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Extracted text:")
	fmt.Println(text)

	meta, err := goka.ExtractMetadataFromFile(test_file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Metadata:", meta)
}
