# Free latin fonts

This project exposes fonts (latin-ext only) in a similar way to the [go-fonts](https://github.com/go-fonts) repositories, but using `embed` and transparently loading the `*sfnt.Font` right away.

This project was made mostly for personal use, to make it easier to write code examples that use the [etxt](https://github.com/tinne26/etxt) text rendering package (Golang / [Ebitengine](https://github.com/hajimehoshi/ebiten)).

## Licenses

Fonts have each their own licenses, located in their respective folders, but they are 
all licensed under [free software licenses](https://en.wikipedia.org/wiki/Free_software).

The code is distributed under the MIT license instead.

## Usage

All subpackages expose three methods:
- `Font()`, which parses the font if it wasn't parsed yet, caches it and returns it as an `*sfnt.Font`.
- `Name()`, which exposes the font name that you would use with `etxt.FontLibrary`.
- `Release()`, which frees the cached font if it was loaded at any point by `Font()`.

Example program (should be run with `-tags gtxt`):
```Golang
package main

import "fmt"
import "github.com/tinne26/etxt"
import "github.com/tinne26/fonts/liberation/lbrtserif"

func main() {
	renderer := etxt.NewStdRenderer()
	renderer.SetFont(lbrtserif.Font())
	rect := renderer.SelectionRect("controversial opinion")
	fmt.Printf("%dx%d pixels\n", rect.Width.Ceil(), rect.Height.Ceil())
}
```
The names of the subpackages and their paths match what you can find in this repository.

## Subpackage cheatsheet

Font names are shortened in funny ways, but they get too long and repetitive for my taste otherwise:
- **Liberation Sans** | *sans-serif, regular, neutral* | `github.com/tinne26/fonts/liberation/lbrtsans`
- **Liberation Serif** | *serif, regular, neutral* | `github.com/tinne26/fonts/liberation/lbrtserif`
- **Liberation Mono** | *monospace, regular, neutral* | `github.com/tinne26/fonts/liberation/lbrtmono`
