package main

import (
	"fmt"
	"github.com/undo-k/datamosh/internal/headers"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func decimate(input []byte, startIndex uint, endIndex uint) []byte {
	data := input[:]
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for idx := startIndex; idx < endIndex; idx++ {
		if r.Float64()*1000000 < 10.0 {
			b := data[idx]
			data[idx] = ((b ^ b) | (0x69 ^ b)) ^ uint8(r.Float32()*10)
		}
	}

	return data
}

func quadratic(input []byte, startIndex uint, endIndex uint) []byte {
	data := input[:]
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for idx := startIndex; idx < endIndex; idx++ {
		if r.Float64()*1000000 < 10.0 {
			data[idx] = (uint8(idx) * uint8(idx)) + (uint8(idx) * data[idx]) + data[idx]
		}
	}

	return data
}

func rotateSubslice(input []byte, startIndex uint, endIndex uint, displacement int) []byte {
	data := input[:]
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	startIndex = uint(r.Intn(int(endIndex-startIndex)) + int(startIndex))
	endIndex = uint(r.Intn(int(endIndex-startIndex)) + int(startIndex))

	sub := data[startIndex:endIndex]

	displacement = displacement % len(sub)

	fmt.Printf("displacement is %v\n", displacement)

	rotatedSlice := append(sub[len(sub)-displacement:], sub[:len(sub)-displacement]...)
	return append(append(data[:startIndex], rotatedSlice...), data[endIndex:]...)
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

	data = decimate(data, start, end)

	data = rotateSubslice(data, start, end, 210000)

	data = quadratic(data, start, end)

	outfile, err := os.Create("output.jpg")
	checkError(err)
	defer outfile.Close()

	outfile.Write(data)

	//fmt.Printf("counted this many byte markers: %v", count)

	//fmt.Println(data)

	//fmt.Println(len(os.Args), os.Args)
}
