package headers

import (
	"errors"
	"log"
	"reflect"
)

func GetFileType(data []byte) string {
	jpegHeader := []byte{0xff, 0xd8}            // https://web.archive.org/web/20120403212223/http://class.ee.iastate.edu/ee528/Reading%20material/JPEG_File_Format.pdf
	pngHeader := []byte{0x89, 0x50, 0x4e, 0x47} // https://en.wikipedia.org/wiki/PNG#File_header
	bmpHeader := []byte{0x42, 0x4D}             // https://en.wikipedia.org/wiki/BMP_file_format#Bitmap_file_header
	switch {
	case reflect.DeepEqual(data[:2], jpegHeader):
		return ".jpg"
	case reflect.DeepEqual(data[:4], pngHeader):
		return ".png"
	case reflect.DeepEqual(data[:2], bmpHeader):
		log.Println("found a bmp hedaer")
		return ".bmp"
	}

	log.Fatal(errors.New("Could not get type from header"))
	return "oops! all out of headers!"
}

func GetBounds(data []byte, header string) (uint, uint) {
	switch header {
	case ".jpg":
		return getJpegBounds(data)
	case ".png":
		return getPngBounds(data)
	case ".bmp":
		return getBmpBounds(data)
	}

	return 0, 0
}
