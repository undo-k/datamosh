package fileio

import (
	"io"
	"log"
	"os"
)

func ReadData(filename *string) []byte {
	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func WriteData(data *[]byte, filename *string) {
	outfile, err := os.Create(*filename)
	if err != nil {
		log.Fatal(err)
	}
	defer outfile.Close()

	outfile.Write(*data)
}
