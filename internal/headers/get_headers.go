package headers

import (
	"os"
	"reflect"
)

func GetFileType(data []byte) string {
	jpegHeader := []byte{0xff, 0xd8}

	if reflect.DeepEqual(data[:2], jpegHeader) {
		return ".jpg"
	}
	// TODO: add more file headerz

	return "oops! all out of headers!"
}

func GetBounds(data []byte, header string) (uint, uint) {
	jpegHeader := []byte{0xff, 0xd8}

	switch header {
	case ".jpg":
		return getJpegBounds(data)
	}

	if reflect.DeepEqual(data[:2], jpegHeader) {
		//fmt.Println("Jpeg header found")
		return getJpegBounds(data)
	}
	os.Exit(1)

	return 0, 0
}
