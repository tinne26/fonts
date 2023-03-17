package lbrtmono

import "golang.org/x/image/font/sfnt"
import "testing"

func TestFont(t *testing.T) {
	if cachedFont != nil { t.Fatal("cachedFont != nil") }
	font := Font()
	if cachedFont == nil { t.Fatal("cachedFont == nil") }
	value, err := font.Name(nil, sfnt.NameIDFamily)
	if err != nil { t.Fatalf("unexpected error: %s", err) }
	name := Name()
	if value != name {
		t.Fatalf("name mismatch: package says '%s', font data says '%s'", name, value)
	}
	Release()
	if cachedFont != nil { t.Fatal("font release failed") }
}
