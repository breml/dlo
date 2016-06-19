package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	d, err := search(os.Args)
	if err != nil {
		log.Fatalf("search failed: %s", err)
	}
	fmt.Print(d)
}

func search(args []string) (*Doc, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("Minimum 1 arg required")
	}

	// Arguments
	// rmWords=off -> entfernt Words aus der XML-Antwort
	// sectLenMax=10 -> search results per section
	resp, err := http.Get("http://dict.leo.org/dictQuery/m-vocab/ende/query.xml?tolerMode=nof&lp=ende&lang=de&rmWords=off&rmSearch=on&search=" + os.Args[1] + "&searchLoc=0&resultOrder=basic&multiwordShowSingle=on&pos=0&sectLenMax=10&n=1")
	if err != nil {
		log.Fatalf("http get error: %s", err)
	}
	defer resp.Body.Close()

	return ProcessQueryXml(args[1], resp.Body)
}
