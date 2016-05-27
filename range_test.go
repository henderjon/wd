package main

import (
	"regexp"
	"testing"
)

func Test_Regex(t *testing.T) {
	var (
		str      string
		expect   int
		matches  [][]string
		examples []string
	)

	// ^(?:(\d+)?,?(\d+|\${1})?)?((?i)[a-z]{1})$
	// ^(?:(\d+)?(?:(?:,)?|(?:,(\d+|\${1})?)))?((?i)[a-z]{1})$
	// ^(?:(\d+)?(?:,?(\d+|\${1})?))?((?i)[a-z]{1})$
	regex := regexp.MustCompile(commandRegex)

	examples = []string{
		`a`,
		`12,34a`,
		`12,a`,
		`,34a`,
		`,$a`,
		`1,$a`,
		`2a`,
		`,a`,
		`2,a`,
	}

	expect = 4 // we should only have 3 captures (+ full string as 0)
	for _, str = range examples {
		matches = regex.FindAllStringSubmatch(str, -1)
		if matches != nil {
			if len(matches[0]) != expect {
				t.Error("match error: \nexpected", expect, "\nactual", len(matches[0]), "\n", str, "=", matches)
			}
		} else {
			t.Error("matches for `", str, "`was nil")
		}
	}

	// these should be nil
	examples = []string{
		`1,2`,
		`1,2ab`,
		`a,2a`,
		`2,$`,
		`2,`,
	}

	for _, str = range examples {
		matches = regex.FindAllStringSubmatch(str, -1)
		if matches != nil {
			t.Error("match error: should be nil\n", str)
		}
	}

}

func Test_Range(t *testing.T) {
	var (
		lr *lRange
		e  error
	)

	fh := NewFile()

	lr, e = parseRange([]byte("a"), fh)
	if e != nil {
		t.Error(e)
	}

	if lr.min != 0 && lr.max != 0 {
		t.Error("range error: min ==", lr.min, "max ==", lr.max)
	}

	fh.lines = make([][]byte, 34)
	lr, e = parseRange([]byte("12,34a"), fh)
	if e != nil {
		t.Error(e)
	}

	if lr.min != 12 && lr.max != 34 {
		t.Error("range error: min ==", lr.min, "max ==", lr.max)
	}

	lr, e = parseRange([]byte(",34a"), fh)
	if e != nil {
		t.Error(e)
	}

	if lr.min != 0 && lr.max != 34 {
		t.Error("range error: min ==", lr.min, "max ==", lr.max)
	}

	lr, e = parseRange([]byte(",$a"), fh)
	if e != nil {
		t.Error(e)
	}

	if lr.min != 0 && lr.max != 34 {
		t.Error("range error: min ==", lr.min, "max ==", lr.max)
	}

	lr, e = parseRange([]byte("34,12a"), fh)
	if e == nil {
		t.Error("range error: should error")
	}

	lr, e = parseRange([]byte("a,$a"), fh)
	if e == nil {
		t.Error("range error: should error")
	}

}
