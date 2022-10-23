package lbrtserif

import "log"
import _ "embed"
import "github.com/tinne26/etxt"

//go:embed LiberationSerif-Regular.ttf.gz
var gzBytes []byte

var cachedFont *etxt.Font

func GzBytes() []byte { return gzBytes }
func FreeFont() { cachedFont = nil }
func Font() *etxt.Font {
	if cachedFont != nil { return cachedFont }

	var err error
	cachedFont, _, err = etxt.ParseFontBytes(gzBytes)
	if err != nil {
		log.Fatal("tinne26/fonts broken version testing: " + err.Error())
	}
	return cachedFont
}
