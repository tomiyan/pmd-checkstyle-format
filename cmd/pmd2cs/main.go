package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/tomiyan/pmd2cs"
	"io"
	"os"
	"strings"
)

const version = "0.0.7"

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

	xml, err := convertXMLString(result)
	if err != nil {
		return err
	}

	fmt.Fprintln(w, xml)
	return nil
}

func convertXMLString(body *pmd2cs.CheckStyleResult) (string, error) {
	formattedBody, err := xml.MarshalIndent(body, "", "    ")
	if err != nil {
		return "", err
	}
	return strings.Join([]string{xml.Header, string(formattedBody)}, ""), nil
}
