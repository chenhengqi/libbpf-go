package main

import (
	"path"
	"strings"
)

func roundUp(x, y int) int {
	return (x + y - 1) / y * y
}

func isLetter(b byte) bool {
	return false
}

func isDigit(b byte) bool {
	return false
}

func objectName(s string) string {
	fileName := path.Base(s)
	fileName = strings.TrimSuffix(fileName, ".o")
	name := []byte(fileName)
	for i, b := range name {
		if !isLetter(b) && !isDigit(b) && b != '_' {
			name[i] = '_'
		}
	}
	return string(name)
}
