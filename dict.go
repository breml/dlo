package main

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strings"

	"github.com/TreyBastian/colourize"
	"github.com/olekukonko/tablewriter"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Doc struct {
	Search string
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

func ProcessQueryXml(search string, input io.Reader) (*Doc, error) {
	b, err := io.ReadAll(input)
	if err != nil {
		return nil, fmt.Errorf("read resp body error: %s", err)
	}

	var d Doc
	if len(b) > 0 {
		if err := xml.Unmarshal(b, &d); err != nil {
			return nil, fmt.Errorf("xml unmarshal error: %s", err)
		}
	}
	d.Search = search

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
		if _, err := fmt.Fprintln(w, s.SectionName); err != nil {
			return ""
		}
		table.Render()
	}
	if err := w.Flush(); err != nil {
		return ""
	}
	b, err := io.ReadAll(r)
	if err != nil {
		return ""
	}
	caser := cases.Title(language.English)
	titleSearch := caser.String(d.Search)
	fmt.Println(titleSearch)
	str := strings.ReplaceAll(string(b), titleSearch, colourize.Colourize(titleSearch, colourize.Bold))
	return strings.ReplaceAll(str, d.Search, colourize.Colourize(d.Search, colourize.Bold))
}
