# Free latin fonts easy to load into etxt

This project exposes a few fonts (latin-ext only) for use with the [etxt](https://github.com/tinne26/etxt) text rendering package (Golang / Ebitengine). Each font style is available in its own subpackage.

This project takes inspiration from [go-fonts](https://github.com/go-fonts) repositories, which operate in a more generic way.

## Licenses

Fonts have each their own licenses, located in their respective folders, but they are 
all licensed under [free software licenses](https://en.wikipedia.org/wiki/Free_software).

The code is distributed under the MIT license instead.

## Usage

All subpackages expose three methods:
- `GzBytes()`, which gives direct access to the gzipped `[]byte`s of the font `.ttf`s. The `ParseFontBytes()` methods on etxt (>= v0.0.7) can automatically detect the gzip format and parse the font correctly, so you can do that directly or add them to font libraries without issue.
- `Font()`, which parses the font if it wasn't parsed yet, caches it and returns it as an `*etxt.Font`.
- `FreeFont()`, which frees the cached font if it was loaded at any point by `Font()`. For completeness.

Example program (should be run with `-tags gtxt`):
```Golang
package main

import "fmt"
import "github.com/tinne26/etxt"
import "github.com/tinne26/fonts/liberation/lbrtserif"

func main() {
	renderer := etxt.NewStdRenderer()
	renderer.SetFont(lbtserif.Font())
	rect := renderer.SelectionRect("please measure me")
	fmt.Printf("%dx%d pixels", rect.Width.Ceil(), rect.Height.Ceil())
}
```
The names of the subpackages and their paths match what you can find in this repository.

## Key differences with go-fonts

- All subpackages depend on etxt and embed. These dependencies are excessive for what the subpackages
  do, but you should only use this project alongside etxt, in which case no new dependencies will
  actually be triggered.
- The font `.ttf.gz` files are directly embedded instead of hardcoded in the source files.
- The font raw data is only exposed gzipped.
- Font names are shortened in funny ways, but they get too long for my taste otherwise.
- The documentation actually talks a bit about the fonts. While I'm not an expert on typography,
  I felt like sharing my own impressions on the fonts and giving due credit to their designers.

## Subpackage cheatsheet

- **Liberation Sans** | *sans-serif, regular, neutral* | `github.com/tinne26/fonts/liberation/lbrtsans`
- **Liberation Serif** | *serif, regular, neutral* | `github.com/tinne26/fonts/liberation/lbrtserif`
- **Liberation Mono** | *monospaced, regular, neutral* | `github.com/tinne26/fonts/liberation/lbrtmono`
