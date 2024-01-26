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

	// Simple addition of heading text
	docx.AddHeading("Title", 0)
	docx.AddHeading("Heading 1", 1)
	docx.AddHeading("Heading 2", 2)
	docx.AddHeading("Heading 3", 3)
	docx.AddHeading("Heading 4", 4)

	err = docx.SaveTo("headings.docx")
	if err != nil {
		log.Fatal(err)
	}
}
