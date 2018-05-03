package struct_tag

import (
	"testing"
)

func TestEmptyLookup(t *testing.T) {
	testCases := [][2]string{
		{``, ``},
		{`a`, ``},
		{`a:""`, ``},
		{`a:"av"`, `b`},
	}
	for _, tc := range testCases {
		if v, ok := Lookup(tc[0], tc[1]); v != "" || ok {
			t.Errorf("Lookup(%#v, %#v) got unexpected: %#v, %v", tc[0], tc[1], v, ok)
		}
	}
}

func TestNonEmptyLookup(t *testing.T) {
	testCases := [][3]string{
		{`a:""`, `a`, ``},
		{`a:"av"`, `a`, `av`},
		{`a:"av" b:"b v"`, `b`, `b v`},
		{`a:"av"
		b:"b v"`, `b`, `b v`},
	}
	for _, tc := range testCases {
		if v, ok := Lookup(tc[0], tc[1]); v != tc[2] || !ok {
			t.Errorf("Lookup(%#v, %#v) got unexpected: %#v, %v", tc[0], tc[1], v, ok)
		}
	}
}
