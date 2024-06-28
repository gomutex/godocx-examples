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

	err = document.SaveTo("pic.docx")
	if err != nil {
		log.Fatal(err)
	}
}
