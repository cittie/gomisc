package gomisc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToCamelCase(t *testing.T) {
	tests := []struct {
		in        string
		needUpper bool
		out       string
	}{
		{"test_case", true, "TestCase"},
		{"test.case", true, "TestCase"},
		{"test", true, "Test"},
		{"TestCase", true, "TestCase"},
		{" test  case ", true, "TestCase"},
		{"", true, ""},
		{"many_many_words", true, "ManyManyWords"},
		{"AnyKind of_string", true, "AnyKindOfString"},
		{"odd-fix", true, "OddFix"},
		{"numbers2And55with000", true, "Numbers2And55With000"},
		{"ID", true, "ID"},
		{"foo-bar", false, "fooBar"},
		{"TestCase", false, "testCase"},
		{"", false, ""},
		{"AnyKind of_string", false, "anyKindOfString"},
		{"AnyKind.of-string", false, "anyKindOfString"},
		{"ID", false, "iD"},
		{"some string", false, "someString"},
		{" some string", false, "someString"},
	}

	for _, test := range tests {
		assert.Equal(t, test.out, ToCamelCase(test.in, test.needUpper), test)
	}
}
