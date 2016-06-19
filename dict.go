package main

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/olekukonko/tablewriter"
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

func ProcessQueryXml(input io.Reader) (*Doc, error) {
	b, err := ioutil.ReadAll(input)
	if err != nil {
		return nil, fmt.Errorf("read resp body error: %s", err)
	}

	var d Doc
	xml.Unmarshal(b, &d)

	return &d, nil
}

func (d Doc) String() string {
	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)
	r := bufio.NewReader(&buf)
	// for _, s := range d.Sections {
	for i := len(d.Sections) - 1; i >= 0; i-- {
		s := d.Sections[i]
		table := tablewriter.NewWriter(w)
		table.SetAutoWrapText(false)
		table.SetHeader([]string{"En", "De"})
		for _, v := range s.Entries {
			table.Append([]string{v.Sides[0].Word, v.Sides[1].Word})
		}
		fmt.Fprintln(w, s.SectionName)
		table.Render()
	}
	w.Flush()
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return ""
	}
	return string(b)
}
