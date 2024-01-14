package headers

import (
	"fmt"
	"os"
	"reflect"
)

func GetHeaderBounds(data []byte) (uint, uint) {
	jpegHeader := []byte{0xff, 0xd8}

	if reflect.DeepEqual(data[:2], jpegHeader) {
		fmt.Println("Jpeg header found")
		return getJpegHeaders(data)
	}
	os.Exit(1)

	return 0, 0
}
