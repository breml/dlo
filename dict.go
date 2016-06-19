package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
)

type Doc struct {
	// 	SectionList []SectionList `xml:"sectionlist"`
	// }
	// type SectionList struct {
	Sections []Section `xml:"sectionlist>section"`
}
type Section struct {
	SectionName string  `xml:"sctName,attr"`
	Entries     []Entry `xml:"entry"`
}
type Entry struct {
	Sides []Side `xml:"side"`
}
type Side struct {
	Lang string `xml:"lang,attr"`
	// 	Words []Word `xml:"words"`
	// }
	// type Word struct {
	Word string `xml:"words>word"`
}

func ProcessQueryXml(input io.Reader) (Doc, error) {
	b, err := ioutil.ReadAll(input)
	if err != nil {
		return Doc{}, fmt.Errorf("read resp body error: %s", err)
	}

	var d Doc
	xml.Unmarshal(b, &d)

	return d, nil
}
