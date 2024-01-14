package moshers

import (
	"fmt"
	"math/rand"
	"time"
)

func Decimate(input []byte, startIndex uint, endIndex uint) []byte {
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

func Quadratic(input []byte, startIndex uint, endIndex uint) []byte {
	data := input[:]
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for idx := startIndex; idx < endIndex; idx++ {
		if r.Float64()*1000000 < 10.0 {
			data[idx] = (uint8(idx) * uint8(idx)) + (uint8(idx) * data[idx]) + data[idx]
		}
	}

	return data
}

func RotateSubslice(input []byte, startIndex uint, endIndex uint, displacement int) []byte {
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