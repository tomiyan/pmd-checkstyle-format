package main

import (
	"flag"
	"fmt"
	"os"
	format "github.com/tomiyan/pmd2cs"
	"encoding/xml"
)

var version = "0.0.3"

func main() {
	var showVersion bool
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&showVersion ,"version", false, "show version")
	flag.Parse()
	if showVersion {
		fmt.Println("version:", version)
		return
	}

	result, err := format.PmdParser{}.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "pmd2cs: %v\n", err)
		os.Exit(1)
		return
	}

	var data []byte
	data = append(data, []byte(xml.Header)...)
	xmlData, err := xml.MarshalIndent(result, "", "    ")

	data = append(data, xmlData...)
	fmt.Println(string(data))
}