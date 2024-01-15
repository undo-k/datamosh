# datamosh
### a commandline tool for datamoshing images
Datamoshing is a sort of glitch art that involves breaking/corrupting filetypes in such a way they're still viewable using standard programs but have distorted the media. Currently supports jpg, png, bmp, and gif files.
### build
`go build -o datamosh.exe .\cmd\datamosh\main.go`
### usage
`./datamosh -i <input file / mandatory> [-o output file (optional)] [shift|decimate|quadtratic (at least one of these)]`

## examples
before and after decimate:

![alt text](https://github.com/undo-k/datamosh/blob/master/examples/input.png?raw=true)
![img.png](https://github.com/undo-k/datamosh/blob/master/examples/output_decimate.png?raw=true)

before and after shift:

![alt text](https://github.com/undo-k/datamosh/blob/master/examples/input_1.jpg?raw=true)
![img.png](https://github.com/undo-k/datamosh/blob/master/examples/output_shift_3000.jpg?raw=true)

before and after quadratic:

![alt text](https://github.com/undo-k/datamosh/blob/master/examples/input.gif?raw=true)
![alt text](https://github.com/undo-k/datamosh/blob/master/examples/output.gif?raw=true)