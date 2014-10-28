package bits

func HammingDistance32(x, y uint32) uint32 {
	return PopCount32(x ^ y)
}

func HammingDistance64(x, y uint64) uint64 {
	return PopCount64(x ^ y)
}
