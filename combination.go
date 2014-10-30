package bits

const (
	Uint64AllZero = uint64(0)
	Uint64AllMask = ^Uint64AllZero
	Mask01        = uint64(0x1)
	Mask11        = uint64(0x3)
)

func CombinationLTEK(n, k uint64) []uint64 {
	var combination []uint64
	for i := uint64(1); i <= k; i++ {
		combination = append(combination, Combination64(n, i)...)
	}
	return combination
}

func Combination64(n, k uint64) []uint64 {
	count := c[n][k]
	combination := make([]uint64, count)

	v := ^(Uint64AllMask << k)
	combination[0] = v

	for i := uint64(1); i < count; i++ {
		for s := uint64(0); s < n-1; s++ {
			mask := Mask11 << s
			if (v&mask)^(Mask01<<s) == 0 {
				v ^= mask
				oneCount := PopCount64(v & ^(Uint64AllMask << s))
				v &= Uint64AllMask << s
				v |= ^(Uint64AllMask << oneCount)
				combination[i] = v
				break
			}
		}
	}
	return combination
}

var c [65][65]uint64

func init() {

	for k := 0; k <= 64; k++ {
		c[0][k] = 0
	}
	for n := 1; n <= 64; n++ {
		c[n][0] = 1
	}

	for n := 1; n <= 64; n++ {
		for k := 1; k <= n; k++ {
			if k == 1 || k == n-1 {
				c[n][k] = uint64(n)
			} else if n == k {
				c[n][k] = 1
			} else {
				c[n][k] = c[n-1][k] + c[n-1][k-1]
			}
		}
	}
}

func CombinationCount(n, k uint64) uint64 {
	return c[n][k]
}
