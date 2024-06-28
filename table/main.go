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

	table := document.AddTable()
	// Predefined style in the document template
	table.Style("LightList-Accent2")

	tblRow := table.AddRow()
	cell00 := tblRow.AddCell()
	cell00.AddParagraph("Column1")
	cell01 := tblRow.AddCell()
	cell01.AddParagraph("Column2")

	tblRow1 := table.AddRow()
	cell10 := tblRow1.AddCell()
	cell10.AddParagraph("Row2 - Column 1")

	cell11 := tblRow1.AddCell()
	cell11.AddParagraph("Row2 - Column 2")

	err = document.SaveTo("table.docx")
	if err != nil {
		log.Fatal(err)
	}
}
