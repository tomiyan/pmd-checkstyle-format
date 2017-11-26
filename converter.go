package pmd_checkstyle_format

import (
	"encoding/xml"
	"io"
	"strings"
	"time"
	"fmt"
)

// PmdStyleParser.
type PmdParser struct{}

func (PmdParser) Parse(r io.Reader) (error, *CheckStyleResult) {
	var decoder = new(PmdResult)
	if err := xml.NewDecoder(r).Decode(decoder); err != nil {
		return err, nil
	}
	var csFiles []*CheckStyleFile
	for _, file := range decoder.Files {
		var csErrors []*CheckStyleError
		for _, violation := range file.Violations {
			var severity string
			switch violation.Priority {
			case 1, 2:
				severity = "error"
			case 3, 4:
				severity = "warning"
			default:
				severity = "info"
			}
			message := fmt.Sprintf("rule: %s, message: %s, externalInfoUrl: %s", violation.Rule, strings.TrimSpace(violation.Message), violation.ExternalInfoUrl)
			csError := &CheckStyleError{
				Line:    violation.BeginLine,
				Message: message,
				Severity: severity,
			}
			csErrors = append(csErrors, csError)
		}
		csFile := &CheckStyleFile{
			Name:   file.Name,
			Errors: csErrors,
		}
		csFiles = append(csFiles, csFile)
	}
	csResult := &CheckStyleResult{
		Files: csFiles,
	}
	return nil, csResult
}

// CheckStyleResult represents pmd XML result.
// <?xml version="1.0" encoding="UTF-8" ?><pmd version="@project.version@" timestamp="2017-11-26T06:00:10+00:00"><file ...></file>...</pmd>
//
// References:
//   - https://phpmd.org
type PmdResult struct {
	XMLName   xml.Name   `xml:"pmd"`
	Version   string     `xml:"version,attr"`
	Timestamp time.Time  `xml:"timestamp,attr"`
	Files     []*PmdFile `xml:"file,omitempty"`
}

// CheckStyleFile represents <file name="fname"><error ... />...</file>
type PmdFile struct {
	Name       string          `xml:"name,attr"`
	Violations []*PmdViolation `xml:"violation"`
}

// CheckStyleError represents <violation beginline="8" endline="8" rule="UnusedLocalVariable" ruleset="Unused Code Rules" externalInfoUrl="http://phpmd.org/rules/unusedcode.html#unusedlocalvariable" priority="3"> Avoid unused local variables such as '$hoge'.</violation>
type PmdViolation struct {
	BeginLine       int    `xml:"beginline,attr,omitempty"`
	EndLine         int    `xml:"endline,attr,omitempty"`
	Rule            string `xml:"rule,attr,omitempty"`
	Ruleset         string `xml:"ruleset,attr,omitempty"`
	ExternalInfoUrl string `xml:"externalInfoUrl,attr,omitempty"`
	Priority        int    `xml:"priority,attr"`
	Package         string `xml:"package,attr,omitempty"`
	Class           string `xml:"class,attr,omitempty"`
	Method          string `xml:"method,attr,omitempty"`
	Message         string `xml:",chardata"`
}

// CheckStyleResult represents checkstyle XML result.
// <?xml version="1.0" encoding="utf-8"?><checkstyle version="4.3"><file ...></file>...</checkstyle>
//
// References:
//   - http://checkstyle.sourceforge.net/
//   - http://eslint.org/docs/user-guide/formatters/#checkstyle
type CheckStyleResult struct {
	XMLName xml.Name          `xml:"checkstyle"`
	Files   []*CheckStyleFile `xml:"file,omitempty"`
}

// CheckStyleFile represents <file name="fname"><error ... />...</file>
type CheckStyleFile struct {
	Name   string             `xml:"name,attr"`
	Errors []*CheckStyleError `xml:"error"`
}

// CheckStyleError represents <error line="1" column="10" severity="error" message="msg" source="src" />
type CheckStyleError struct {
	Column   int    `xml:"column,attr,omitempty"`
	Line     int    `xml:"line,attr"`
	Message  string `xml:"message,attr"`
	Severity string `xml:"severity,attr,omitempty"`
	Source   string `xml:"source,attr,omitempty"`
}
