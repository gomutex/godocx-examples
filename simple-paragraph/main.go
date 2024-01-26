package main

import (
	"log"

	"github.com/gomutex/godocx"
)

func main() {
	docx, err := godocx.NewDocument()
	if err != nil {
		log.Fatal(err)
	}

	// Simple addition of paragraph text
	docx.AddParagraph("Hello, World")

	// Add empty paragraph elemenet then later add text with run
	p := docx.AddEmptyParagraph()
	// Using Run
	r := p.AddText("Hello, Parallel World")
	r.Color("bc00d6")

	err = docx.SaveTo("simple-para.docx")
	if err != nil {
		log.Fatal(err)
	}
}
