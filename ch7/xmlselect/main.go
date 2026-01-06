// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 214.
//!+

// Xmlselect prints the text of selected elements of an XML document.
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

type Tag struct {
	name, id string
	class    []string
}

func parseTag(raw string) *Tag {
	var name string
	id := ""
	class := make([]string, 0)
	idIdx := strings.Index(raw, "#")
	classIdx := strings.Index(raw, ".")
	if idIdx == -1 && classIdx == -1 {
		return &Tag{raw, id, class}
	}
	if idIdx != -1 {
		name = raw[:idIdx]
		if classIdx == -1 {
			id = raw[idIdx+1:]
		} else {
			id = raw[idIdx+1 : classIdx]
		}

	}
	if classIdx != -1 {
		if idIdx == -1 {
			name = raw[:classIdx]
		}
		raw = raw[classIdx+1:]
		for classIdx := strings.Index(raw, "."); classIdx != -1; classIdx = strings.Index(raw, ".") {
			class = append(class, raw[:classIdx])
			raw = raw[classIdx+1:]
		}
		class = append(class, raw)
	}
	return &Tag{name, id, class}
}

func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []xml.StartElement // stack of element names
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok) // push
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			selected := make([]*Tag, 0)
			for i := 1; i < len(os.Args); i++ {
				selected = append(selected, parseTag(os.Args[i]))
			}
			if containsAll(stack, selected) {
				names := make([]string, len(stack))
				for i, el := range stack {
					names[i] = el.Name.Local
				}
				fmt.Printf("%s: %s\n", strings.Join(names, " "), tok)
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x []xml.StartElement, y []*Tag) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if contains(x[0], y[0]) {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

func contains(xl xml.StartElement, yl *Tag) bool {
	if xl.Name.Local != yl.name {
		return false
	}

	var elemID string
	var elemClassStr string
	for _, attr := range xl.Attr {
		if attr.Name.Local == "id" {
			elemID = attr.Value
		} else if attr.Name.Local == "class" {
			elemClassStr = attr.Value
		}
	}

	if yl.id != "" && yl.id != elemID {
		return false
	}

	if len(yl.class) != 0 {
		attrClasses := strings.Split(elemClassStr, " ")
		for _, classEl := range yl.class {
			if !slices.Contains(attrClasses, classEl) {
				return false
			}
		}
	}

	return true
}

//!-
