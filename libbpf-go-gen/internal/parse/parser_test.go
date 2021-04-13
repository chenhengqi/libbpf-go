package parse

import (
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	enums := map[string]struct{}{}
	structs := map[string]struct{}{}
	for _, testcase := range strings.Split(testdata, "\n") {
		parser := NewParser([]byte(testcase))
		types := parser.Parse()
		if len(types) == 0 {
			t.Fatalf("failed to parse `%s`", testcase)
		}
		for _, typeSpec := range types {
			if typeSpec.Name == "" {
				if typeSpec.IsVoid {
					// int foo(void)
					t.Logf("no params: %s", testcase)
				} else {
					t.Fatalf("name NOT set: %s", testcase)
				}
			}
			if typeSpec.TypeName == "" {
				t.Fatalf("type name NOT set: `%s`", testcase)
			}
			if typeSpec.IsEnum {
				enums[typeSpec.TypeName] = struct{}{}
			}
			if typeSpec.IsStruct {
				structs[typeSpec.TypeName] = struct{}{}
			}
		}
	}
	t.Log("enums:", len(enums), enums)
	t.Log("structs:", len(structs), structs)
}
