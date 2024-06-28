package main

import (
	"log"

	"github.com/gomutex/godocx"
	"github.com/gomutex/godocx/wml/stypes"
)

func main() {
	document, err := godocx.NewDocument()
	if err != nil {
		log.Fatal(err)
	}

	// Simple addition of paragraph text
	p := document.AddParagraph("Hello, World.")
	r := p.AddText(" Italic")
	r.Italic(true)

	// Add empty paragraph elemenet then later add text with run
	p1 := document.AddEmptyParagraph()
	// Using Run
	r1 := p1.AddText("Hello, Parallel World")
	r1.Color("bc00d6")
	r1.Bold(true)

	document.AddEmptyParagraph().AddText("Strike").Strike(true)
	document.AddEmptyParagraph().AddText("Underline").Underline(stypes.UnderlineSingle)
	document.AddEmptyParagraph().AddText("Highlight").Highlight(stypes.ColorIndexBlue)
	document.AddEmptyParagraph().AddText("Shading").Shading(stypes.ShdSolid, "auto", "FF00A0")

	jp1 := document.AddParagraph("Center Justified")
	jp1.Justification("center")

	err = document.SaveTo("paragraph.docx")
	if err != nil {
		log.Fatal(err)
	}
}
