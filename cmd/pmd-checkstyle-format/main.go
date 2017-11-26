package main

import (
	"flag"
	"fmt"
	"os"
	format "github.com/tomiyan/pmd-checkstyle-format"
	"encoding/xml"
)

var version = "0.0.1"

func main() {
	var showVersion bool
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&showVersion ,"version", false, "show version")
	flag.Parse()
	if showVersion {
		fmt.Println("version:", version)
		return
	}

	err, result := format.PmdParser{}.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "pmd-checkstyle-format: %v\n", err)
		os.Exit(1)
		return
	}

	var data []byte
	data = append(data, []byte(xml.Header)...)
	xmlData, err := xml.MarshalIndent(result, "", "    ")

	data = append(data, xmlData...)
	fmt.Println(string(data))
}