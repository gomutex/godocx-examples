package main

import (
	"log"

	"github.com/gomutex/godocx"
)

func main() {
	document, err := godocx.NewDocument()
	if err != nil {
		log.Fatal(err)
	}

	// Simple addition of heading text
	document.AddHeading("Title", 0)
	document.AddHeading("Heading 1", 1)
	document.AddHeading("Heading 2", 2)
	document.AddHeading("Heading 3", 3)
	document.AddHeading("Heading 4", 4)

	err = document.SaveTo("headings.docx")
	if err != nil {
		log.Fatal(err)
	}
}
