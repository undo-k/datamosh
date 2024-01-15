package headers

import (
	"fmt"
	"os"
	"reflect"
)

func GetFileType(data []byte) string {
	jpegHeader := []byte{0xff, 0xd8}            // https://web.archive.org/web/20120403212223/http://class.ee.iastate.edu/ee528/Reading%20material/JPEG_File_Format.pdf
	pngHeader := []byte{0x89, 0x50, 0x4e, 0x47} // https://en.wikipedia.org/wiki/PNG#File_header

	if reflect.DeepEqual(data[:2], jpegHeader) {
		return ".jpg"
	}
	// TODO: add more file headerz

	if reflect.DeepEqual(data[:4], pngHeader) {
		fmt.Println("png header found!")
		return ".png"
	}

	return "oops! all out of headers!"
}

func GetBounds(data []byte, header string) (uint, uint) {
	jpegHeader := []byte{0xff, 0xd8}

	switch header {
	case ".jpg":
		return getJpegBounds(data)
	case ".png":
		return getPngBounds(data)
	}

	if reflect.DeepEqual(data[:2], jpegHeader) {
		//fmt.Println("Jpeg header found")
		return getJpegBounds(data)
	}
	os.Exit(1)

	return 0, 0
}
