# datamosh
### a commandline tool for datamoshing images
Datamoshing is a sort of glitch art that involves breaking/corrupting filetypes in such a way they're still viewable using standard programs but have distorted the media. Currently supports jpg, png, bmp, and gif files.
### build
`go build -o datamosh.exe .\cmd\datamosh\main.go`
### usage
`./datamosh -i <input file / mandatory> [-o output file (optional)] [shift|decimate|quadtratic|blend|add (at least one of these)]`

## examples
quadratic mosh on png:

![quadratic before and after](https://github.com/undo-k/datamosh/blob/master/examples/before_after_quadratic.png?raw=true)

before and after shift on jpg:

![shift before and after](https://github.com/undo-k/datamosh/blob/master/examples/before_after_shift.png?raw=true)

blending several gifs:

![blended gifs](https://github.com/undo-k/datamosh/blob/master/examples/output_blend.GIF?raw=true)
