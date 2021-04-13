package parse

import (
	"bytes"
	"io"
)

// Tokenizer is used to tokenize function signature
type Tokenizer struct {
	src   []byte
	index int
}

// NewTokenizer creates a new Tokenizer
func NewTokenizer(src []byte) *Tokenizer {
	return &Tokenizer{
		src:   src,
		index: 0,
	}
}

func (t *Tokenizer) peek() (byte, error) {
	if t.index >= len(t.src) {
		return byte(0), io.EOF
	}
	return t.src[t.index], nil
}

// Next returns tokens one by one
func (t *Tokenizer) Next() (Token, string, error) {
scan:
	b, err := t.peek()
	if err != nil {
		return EOF, "", err
	}
	switch b {
	case ' ', '\t', '\n', '\r':
		t.index++
		goto scan
	case '(':
		t.index++
		return LeftParenthesis, "(", nil
	case ')':
		t.index++
		return RightParenthesis, ")", nil
	case ',':
		t.index++
		return Comma, ",", nil
	case '*':
		t.index++
		return Asterisk, "*", nil
	default:
		if 'a' <= b && b <= 'z' {
			return t.readSymbol()
		}
		if 'A' <= b && b <= 'Z' {
			return t.readSymbol()
		}
		if b <= '_' {
			return t.readSymbol()
		}
		panic("should not reach here")
	}
}

func (t *Tokenizer) readSymbol() (Token, string, error) {
	sym := bytes.NewBuffer(nil)
	for t.index < len(t.src) {
		if ('a' <= t.src[t.index] && t.src[t.index] <= 'z') ||
			('A' <= t.src[t.index] && t.src[t.index] <= 'Z') ||
			('0' <= t.src[t.index] && t.src[t.index] <= '9') ||
			t.src[t.index] == '_' {
			sym.WriteByte(t.src[t.index])
			t.index++
		} else {
			break
		}
	}
	// TODO: currently we return `Type` and `Ident` here
	//       but we should only return `Ident` and distinguish
	//       `Type` and `Ident` at Parser phase
	return toToken(sym.String()), sym.String(), nil
}
