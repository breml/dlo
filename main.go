package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/olekukonko/tablewriter"
)

func main() {
	// Arguments
	// rmWords=off -> entfernt Words aus der XML-Antwort
	resp, err := http.Get("http://dict.leo.org/dictQuery/m-vocab/ende/query.xml?tolerMode=nof&lp=ende&lang=de&rmWords=off&rmSearch=on&search=" + os.Args[1] + "&searchLoc=0&resultOrder=basic&multiwordShowSingle=on&pos=0&sectLenMax=16&n=1")
	if err != nil {
		log.Fatalf("http get error: %s", err)
	}
	defer resp.Body.Close()

	d, err := ProcessQueryXml(resp.Body)

	// fmt.Println(d)

	// for _, s := range d.Sections {
	for i := len(d.Sections) - 1; i >= 0; i-- {
		s := d.Sections[i]
		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoWrapText(false)
		table.SetHeader([]string{"En", "De"})
		for _, v := range s.Entries {
			table.Append([]string{v.Sides[0].Word, v.Sides[1].Word})
		}
		fmt.Println(s.SectionName)
		table.Render() // Send output
	}
}
