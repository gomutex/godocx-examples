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

	document.AddHeading("Document Title", 0)

	// Simple addition of paragraph text
	document.AddParagraph("Hello, World")

	document.AddHeading("Heading, level 1", 1)

	// Add empty paragraph elemenet then later add text with run
	p := document.AddEmptyParagraph()
	// Using Run
	r := p.AddText("Hello, Parallel World")
	r.Color("bc00d6")
	r.Size(18)

	p1 := document.AddParagraph("Intense Quote")
	p1.Style("IntenseQuote")

	p = document.AddParagraph("A plain paragraph having some ")
	p.AddText("bold").Bold(true)
	p.AddText(" and some ")
	p.AddText("italic.").Italic(true)

	bulletPara := document.AddParagraph("first item in unordered list")
	bulletPara.Style("List Bullet")

	numberPara := document.AddParagraph("first item in ordered list")
	numberPara.Style("List Number")

	err = document.SaveTo("hello-world.docx")
	if err != nil {
		log.Fatal(err)
	}
}
