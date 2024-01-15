package headers

func getPngBounds(data []byte) (uint, uint) {
	return 10, uint(len(data) - 1)
}
