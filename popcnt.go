package bits

func PopCount32(x uint32) uint32 {
	x -= (x >> 1) & 0x55555555
	x = (x>>2)&0x33333333 + x&0x33333333
	x += x >> 4
	x &= 0x0f0f0f0f
	x *= 0x01010101
	return x >> 24
}

func PopCount64(x uint64) uint64 {
	x -= (x >> 1) & 0x5555555555555555
	x = (x>>2)&0x3333333333333333 + x&0x3333333333333333
	x += x >> 4
	x &= 0x0f0f0f0f0f0f0f0f
	x *= 0x0101010101010101
	return x >> 56
}
