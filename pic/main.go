package main

import (
	"log"

	"github.com/gomutex/godocx"
	"github.com/gomutex/godocx/common/units"
)

func main() {
	document, err := godocx.NewDocument()
	if err != nil {
		log.Fatal(err)
	}

	document.AddPicture("gopher.png", units.Inch(2.9), units.Inch(2.9))

	// Add in existing paragraph
	p := document.AddParagraph("Hello ")
	p.AddPicture("gopher.png", units.Inch(1), units.Inch(1))

	table := document.AddTable()
	// Predefined style in the document template
	table.Style("LightList-Accent3")

	// Insert in Table
	tblRow := table.AddRow()
	cell00 := tblRow.AddCell()
	cell00.AddParagraph("Column1")
	cell01 := tblRow.AddCell()
	cell01.AddParagraph("Column2")

	tblRow1 := table.AddRow()
	cell10 := tblRow1.AddCell()
	cell10.AddParagraph("Row2 - Column 1")
	cell11 := tblRow1.AddCell()
	p2 := cell11.AddEmptyPara()
	p2.AddPicture("gopher.png", units.Inch(1), units.Inch(1))

	err = document.SaveTo("pic.docx")
	if err != nil {
		log.Fatal(err)
	}
}
