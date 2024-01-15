package moshers

import (
	"fmt"
	"math/rand"
	"time"
)

func Decimate(data *[]byte, startIndex uint, endIndex uint, threshold float32) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for idx := startIndex; idx < endIndex; idx++ {
		if r.Float32()*threshold < 1.0 {
			b := uint8(r.Uint32())
			(*data)[idx] = ((b ^ b) | (0x69 ^ b)) ^ uint8(r.Float32()*10)
		}
	}
}

func Quadratic(data *[]byte, startIndex uint, endIndex uint, threshold float32) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for idx := startIndex; idx < endIndex; idx++ {
		if r.Float32()*threshold < 1.0 {
			yIndex := (idx * idx) + (uint((*data)[idx]) * idx) + uint((*data)[idx])
			yIndex = yIndex % endIndex
			if yIndex < startIndex {
				yIndex += startIndex
			}
			temp := (*data)[idx]
			(*data)[idx] = (*data)[yIndex]
			(*data)[yIndex] = temp
		}
	}
}

func ShiftSubslice(data *[]byte, startIndex uint, endIndex uint, displacement int) []byte {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	startIndex = uint(r.Intn(int(endIndex-startIndex)) + int(startIndex))
	endIndex = uint(r.Intn(int(endIndex-startIndex)) + int(startIndex))

	sub := (*data)[startIndex:endIndex]

	displacement = displacement % len(sub)

	fmt.Printf("displacement is %v\n", displacement)

	rotatedSlice := append(sub[len(sub)-displacement:], sub[:len(sub)-displacement]...)
	return append(append((*data)[:startIndex], rotatedSlice...), (*data)[endIndex:]...)
}
