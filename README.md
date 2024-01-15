# datamosh
### a commandline tool for datamoshing images
datamoshing is a sort of glitch art that involves breaking/corrupting filetypes in such a way they're still viewable using standard programs but have distorted the media

### build
`go build -o datamosh.exe .\cmd\datamosh\main.go`
### usage
`./datamosh -i <input file / mandatory> [-o output file (optional)] [shift|decimate|quadtratic (at least one of these)]`