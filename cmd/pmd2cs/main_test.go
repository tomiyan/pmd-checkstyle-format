package main

import (
	"testing"
	"bytes"
	"strings"
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
