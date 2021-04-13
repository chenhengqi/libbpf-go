package parse

import (
	"strings"
	"testing"
)

func TestTokenize(t *testing.T) {
	for _, testcase := range strings.Split(testdata, "\n") {
		tokenizer := NewTokenizer([]byte(testcase))
		for token, literal, err := tokenizer.Next(); err == nil; {
			if token == Type {
				if false {
					t.Logf("type: %s", literal)
				}
			}
			if token == Ident {
				if false {
					t.Logf("ident: %s", literal)
				}
			}
			token, literal, err = tokenizer.Next()
		}
	}
}
