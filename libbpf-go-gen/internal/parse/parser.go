package parse

import (
	"fmt"
	"io"
)

// Parser is used to parse function signature
type Parser struct {
	tokenizer *Tokenizer
	token     Token
	literal   string
	types     []*TypeSpec
}

// NewParser creates a new parser
func NewParser(src []byte) *Parser {
	return &Parser{
		tokenizer: NewTokenizer(src),
	}
}

// Parse parses function arguments
func (p *Parser) Parse() []*TypeSpec {
	for {
		p.next()
		switch p.token {
		case EOF:
			return p.types
		case Const, Struct, Enum, Unsigned, Type:
			p.parseTypes()
		default:
			panic(fmt.Sprintf("unexpected token: %+v, literal: %s", p.token, p.literal))
		}
	}
}

func (p *Parser) next() {
	token, literal, err := p.tokenizer.Next()
	if err != nil {
		if err == io.EOF {
			p.token = EOF
			p.literal = ""
			return
		}
		panic(err)
	}
	p.token = token
	p.literal = literal
}

func (p *Parser) parseTypes() {
	for {
		switch p.token {
		case RightParenthesis:
			return
		case LeftParenthesis, Comma:
			p.next()
			typeSpec := p.parseType()
			p.types = append(p.types, typeSpec)
		default:
			typeSpec := p.parseType()
			p.types = append(p.types, typeSpec)
		}
	}
}

func (p *Parser) parseType() *TypeSpec {
	typeSpec := &TypeSpec{}
	for {
		switch p.token {
		case EOF:
			panic("unexpected EOF")
		case LeftParenthesis, RightParenthesis, Comma:
			return typeSpec
		case Asterisk:
			typeSpec.IsPointer = true
			typeSpec.Indirections++
		case Const:
			typeSpec.IsConst = true
		case Struct:
			typeSpec.IsStruct = true
		case Enum:
			typeSpec.IsEnum = true
		case Unsigned:
			typeSpec.IsUnsigned = true
		case Type:
			if typeSpec.TypeName != "" {
				// we reach here because we have two `Type`
				// currently `btf` and `btf_ext` exist
				// in both `Type` and `Ident`
				if p.literal != "btf" && p.literal != "btf_ext" {
					panic("type name already set: " + typeSpec.TypeName)
				}
				typeSpec.Name = p.literal
				break
			}
			typeSpec.TypeName = p.literal
			if typeSpec.TypeName == "void" {
				typeSpec.IsVoid = true
			}
		case Ident:
			if typeSpec.Name != "" {
				// we reach here because we have two `Ident`
				panic("name already set: " + typeSpec.Name)
			}
			typeSpec.Name = p.literal
		}
		p.next()
	}
}
