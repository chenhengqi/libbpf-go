package parse

type Token int

// tokens
const (
	// void func_name(const struct foo *x, unsigned int n, enum bar z);
	EOF              Token = iota
	LeftParenthesis        // (
	RightParenthesis       // )
	Comma                  // ,
	Asterisk               // *
	Const                  // const
	Struct                 // struct
	Enum                   // enum
	Unsigned               // unsigned
	Type                   // bpf_create_map_attr
	Ident                  // bpf_create_map_xattr
)

var stringToToken = map[string]Token{
	"(":        LeftParenthesis,
	")":        RightParenthesis,
	",":        Comma,
	"*":        Asterisk,
	"const":    Const,
	"struct":   Struct,
	"enum":     Enum,
	"unsigned": Unsigned,
}

var tokenToString = map[Token]string{
	LeftParenthesis:  "(",
	RightParenthesis: ")",
	Comma:            ",",
	Asterisk:         "*",
	Const:            "const",
	Struct:           "struct",
	Enum:             "enum",
	Unsigned:         "unsigned",
	Type:             "type",
	Ident:            "ident",
}

func (t Token) String() string {
	return tokenToString[t]
}

func toToken(s string) Token {
	t, ok := stringToToken[s]
	if ok {
		return t
	}
	_, ok = types[s]
	if ok {
		return Type
	}
	return Ident
}
