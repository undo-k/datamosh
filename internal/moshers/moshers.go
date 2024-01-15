package moshers

import (
	"fmt"
	"math"
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

func Blender(data *[]byte, data2 *[]byte, startIndex uint, endIndex uint, startIndex2 uint, endIndex2 uint, threshold float32, percentage int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	smallImage, bigImage, smallStart, smallEnd, bigStart, bigEnd := minMaxSlice(data, data2, startIndex, endIndex, startIndex2, endIndex2)

	smallStart, smallEnd = randomWindow(smallStart, smallEnd, percentage)

	startIndex = uint(r.Intn(int(endIndex-startIndex)) + int(startIndex))
	endIndex = uint(r.Intn(int(endIndex-startIndex)) + int(startIndex))

	subSmall := (*smallImage)[smallStart:smallEnd]
	subBig := (*bigImage)[bigStart:bigEnd]

	for i := 1; i < int(math.Min(float64(len(subBig)), float64(len(subSmall)))); i++ {
		if r.Float32()*threshold < 1.0 {
			subBig[i] = subSmall[i]
		}
	}

	newBig := append((*bigImage)[:bigStart], subBig...)
	*data = append(newBig, (*bigImage)[bigEnd+1:]...)
}

func Appender(data *[]byte, data2 *[]byte, startIndex uint, endIndex uint, startIndex2 uint, endIndex2 uint, percentage int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	smallImage, bigImage, smallStart, smallEnd, bigStart, bigEnd := minMaxSlice(data, data2, startIndex, endIndex, startIndex2, endIndex2)

	smallStart, smallEnd = randomWindow(smallStart, smallEnd, percentage)
	bigStart, bigEnd = randomWindow(bigStart, bigEnd, percentage)

	startIndex = uint(r.Intn(int(endIndex-startIndex)) + int(startIndex))
	endIndex = uint(r.Intn(int(endIndex-startIndex)) + int(startIndex))

	subSmall := (*smallImage)[smallStart:smallEnd]
	subBig := (*bigImage)[bigStart:bigEnd]

	combo := append(subBig, subSmall...)
	*data = append(*bigImage, combo...)
}

func randomWindow(start uint, end uint, percentage int) (uint, uint) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	percent := float64(percentage) * 0.01
	length := end - start
	windowSize := int(float64(length)*percent) + 1
	newStart := r.Intn(int(end) - windowSize - 1)
	newEnd := newStart + windowSize
	return uint(newStart), uint(newEnd)

}

func minMaxSlice(slice1 *[]byte, slice2 *[]byte, start1 uint, end1 uint, start2 uint, end2 uint) (*[]byte, *[]byte, uint, uint, uint, uint) {
	if len(*slice1) > len(*slice2) {
		return slice2, slice1, start2, end2, start1, end1
	} else {
		return slice1, slice2, start1, end1, start2, end2
	}
}

func clamp(value uint, max uint) uint {
	if value > max {
		return max
	} else {
		return value
	}
}
