package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/tomiyan/pmd2cs"
	"io"
	"os"
)

const version = "0.0.4"

type option struct {
	version bool
}

var opt = &option{}

func init() {
	flag.BoolVar(&opt.version, "v", false, "show version")
	flag.BoolVar(&opt.version, "version", false, "show version")
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
	result, err := pmd2cs.PmdParser{}.Parse(r)
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
