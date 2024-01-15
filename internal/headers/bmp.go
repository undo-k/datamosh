package headers

import (
	"encoding/binary"
)

func getBmpBounds(data []byte) (uint, uint) {
	offsetBytes := data[10:15]

	indexOffset := binary.LittleEndian.Uint32(offsetBytes)

	return uint(indexOffset), uint(len(data) - 1)
}
