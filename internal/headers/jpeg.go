package headers

func getJpegBounds(data []byte) (uint, uint) {
	//var soiIndex uint
	//var appIndex uint
	//var dqtIndex uint
	var sofIndex uint
	//var dhtIndex uint
	//var sosIndex uint
	var eoiIndex uint

	for idx, byteData := range data {
		if byteData == 0xff {
			switch data[idx+1] {
			//case 0xd8:
			//	soiIndex = uint(idx)
			//	fmt.Printf("Found SOI at %v\n", soiIndex)
			//
			//case 0xe0:
			//	appIndex = uint(idx)
			//	fmt.Printf("Found APP0 at %v\n", appIndex)
			//
			//case 0xdb:
			//	dqtIndex = uint(idx)
			//	fmt.Printf("Found DQT at %v\n", dqtIndex)
			case 0xc0:
				sofIndex = uint(idx + 3)
				//fmt.Printf("Found SOF0 at %v\n", sofIndex)
			//case 0xc4:
			//
			//	dhtIndex = uint(idx)
			//	fmt.Printf("Found DHT at %v\n", dhtIndex)
			//case 0xda:
			//
			//	sosIndex = uint(idx + 3)
			//	fmt.Printf("Found SOS at %v\n", sosIndex)
			case 0xd9:

				eoiIndex = uint(idx - 2)
				//fmt.Printf("Found EOI at %v\n", eoiIndex)
			}
		}
	}

	return sofIndex, eoiIndex
}
