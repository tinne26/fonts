package lbrtserif

import _ "embed"
import "golang.org/x/image/font/sfnt"

//go:embed LiberationSerif-Regular.ttf
var bytes []byte

var cachedFont *sfnt.Font

func Release() { cachedFont = nil }
func Name() string { return "Liberation Serif" }
func Font() *sfnt.Font {
	if cachedFont != nil { return cachedFont }
	
	var err error
	cachedFont, err = sfnt.Parse(bytes)
	if err != nil {
		panic("tinne26/fonts broken version testing: " + err.Error())
	}
	return cachedFont
}
