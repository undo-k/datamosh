package headers

import (
	"fmt"
	"math"
)

func getGifBounds(data []byte) (uint, uint) {
	GCTF := data[10]
	fmt.Printf("GCTF bits: %b\n", GCTF)
	GCTFMask := byte(0x0F) // 0b00001111

	GCTSizePow := GCTF & GCTFMask
	fmt.Printf("GCTSizePow: %d\n", GCTSizePow)

	GCTSize := math.Pow(2, float64(GCTSizePow)+1.0)
	fmt.Printf("GCTSize: %d\n", int(GCTSize))

	return uint(GCTSize + 13), uint(len(data) - 1)
}
