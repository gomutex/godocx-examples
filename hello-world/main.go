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

	docx.AddHeading("Document Title", 0)

	// Simple addition of paragraph text
	docx.AddParagraph("Hello, World")

	docx.AddHeading("Heading, level 1", 1)

	// Add empty paragraph elemenet then later add text with run
	p := docx.AddEmptyParagraph()
	// Using Run
	r := p.AddText("Hello, Parallel World")
	r.Color("bc00d6")
	r.Size(18)

	p1 := docx.AddParagraph("Intense Quote")
	p1.Style("IntenseQuote")

	p = docx.AddParagraph("A plain paragraph having some ")
	p.AddText("bold").Bold(true)
	p.AddText(" and some ")
	p.AddText("italic.").Italic(true)

	err = docx.SaveTo("hello-world.docx")
	if err != nil {
		log.Fatal(err)
	}
}
