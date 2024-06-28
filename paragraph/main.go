package main

import (
	"log"

	"github.com/gomutex/godocx"
	"github.com/gomutex/godocx/wml/stypes"
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
	docx.AddEmptyParagraph().AddText("Underline").Underline(stypes.UnderlineSingle)
	docx.AddEmptyParagraph().AddText("Highlight").Highlight(stypes.ColorIndexBlue)
	docx.AddEmptyParagraph().AddText("Shading").Shading(stypes.ShdSolid, "auto", "FF00A0")

	jp1 := docx.AddParagraph("Center Justified")
	jp1.Justification("center")

	err = docx.SaveTo("paragraph.docx")
	if err != nil {
		log.Fatal(err)
	}
}
