package main

import (
	"flag"
	"fmt"
	"os"
	format "github.com/tomiyan/pmd2cs"
	"encoding/xml"
	"io"
)

const version = "0.0.4"

type option struct {
	version bool
}

var opt = &option{}

func init() {
	flag.BoolVar(&opt.version, "v", false, "show version")
	flag.BoolVar(&opt.version ,"version", false, "show version")
}

func main() {
	flag.Parse()

	if err := run(os.Stdin, os.Stdout, opt); err != nil {
		fmt.Fprintf(os.Stderr, "pmd2cs: %v\n", err)
		os.Exit(1)
	}
}

func run(r io.Reader, w io.Writer, o *option) error {
	if o.version {
		fmt.Fprintln(w, version)
		return nil
	}

	result, err := format.PmdParser{}.Parse(r)
	if err != nil {
		return err
	}

	var data []byte
	data = append(data, []byte(xml.Header)...)
	xmlData, err := xml.MarshalIndent(result, "", "    ")

	data = append(data, xmlData...)
	fmt.Fprintln(w, string(data))
	return nil
}