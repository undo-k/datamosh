package main

import (
	"github.com/undo-k/datamosh/internal/headers"
	"github.com/undo-k/datamosh/internal/moshers"
	"io"
	"log"
	"os"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	//if len(os.Args) < 2 {
	//	checkError(errors.New("no command line arguments passed"))
	//}

	filename := "test_image.jpg"

	file, err := os.Open(filename)
	checkError(err)
	defer file.Close()

	data, err := io.ReadAll(file)

	start, end := headers.GetHeaderBounds(data)

	data = moshers.Decimate(data, start, end)

	data = moshers.RotateSubslice(data, start, end, 210000)

	data = moshers.Quadratic(data, start, end)

	outfile, err := os.Create("output.jpg")
	checkError(err)
	defer outfile.Close()

	outfile.Write(data)

	//fmt.Printf("counted this many byte markers: %v", count)

	//fmt.Println(data)

	//fmt.Println(len(os.Args), os.Args)
}
