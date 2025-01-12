package main

import (
	"encoding/xml"
	"os"
	"testing"

	"github.com/gomutex/godocx"
)

func TestMain(t *testing.T) {
	document, err := godocx.OpenDocument("hello-world.docx")
	if err != nil {
		t.Fatal(err)
	}

	xmlEncoder := xml.NewEncoder(os.Stdout)
	err = document.Document.Body.MarshalXML(xmlEncoder, xml.StartElement{})
	if err != nil {
		t.Fatal(err)
	}
}
