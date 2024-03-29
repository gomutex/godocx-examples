package main

import (
	"log"

	"github.com/gomutex/godocx"
	"github.com/gomutex/godocx/oxml/elements"
)

func main() {
	docx, err := godocx.NewDocument()
	if err != nil {
		log.Fatal(err)
	}

	// Simple addition of paragraph text
	p := docx.AddParagraph("Hello, World.")
	r := p.AddText(" Italic")
	r.Italic(true)

	// Add empty paragraph elemenet then later add text with run
	p1 := docx.AddEmptyParagraph()
	// Using Run
	r1 := p1.AddText("Hello, Parallel World")
	r1.Color("bc00d6")
	r1.Bold(true)

	docx.AddEmptyParagraph().AddText("Strike").Strike(true)
	docx.AddEmptyParagraph().AddText("Underline").Underline(elements.UnderlineSingle)
	docx.AddEmptyParagraph().AddText("Highlight").Highlight(elements.ColorIndexBlue)
	docx.AddEmptyParagraph().AddText("Shading").Shading(elements.ShadingTypeSolid, "auto", "FF00A0")

	err = docx.SaveTo("paragraph.docx")
	if err != nil {
		log.Fatal(err)
	}
}
