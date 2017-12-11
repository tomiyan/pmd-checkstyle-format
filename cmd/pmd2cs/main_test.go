package main

import (
	"bytes"
	"strings"
	"testing"
	"github.com/tomiyan/pmd2cs"
	"encoding/xml"
)

func TestRun_version(t *testing.T) {
	stdout := new(bytes.Buffer)
	opt := &option{
		version: true,
	}
	if err := run(nil, stdout, opt); err != nil {
		t.Error(err)
	}
	if got := strings.TrimRight(stdout.String(), "\n"); got != version {
		t.Errorf("version = %v, want %v", got, version)
	}
}

func TestRun_eof(t *testing.T) {
	stdout := new(bytes.Buffer)
	if err := run(new(bytes.Buffer), stdout, opt); err != nil {
		t.Error(err)
	}
	if stdout.String() != "" {
		t.Error("not empty")
	}
}

func TestConvertXMLString(t *testing.T) {
	csr := &pmd2cs.CheckStyleResult{
		XMLName: xml.Name{},
	}
	got, err := convertXMLString(csr)
	if err != nil {
		t.Error(err)
	}
	expected := strings.Join([]string{xml.Header, "<checkstyle></checkstyle>"}, "")
	if got != expected {
		t.Errorf("got %v, want %v", got, expected)
	}
}