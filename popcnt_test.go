package bits

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func LinerPopCount64(x uint64) (n uint64) {
	for i := uint64(0); i < 64; i++ {
		n += (x >> i) & 0x1
	}
	return n
}

func TestPopCount64(t *testing.T) {
	assert.Equal(t, PopCount64(0), 0)
	assert.Equal(t, PopCount64(^uint64(0)), 64)
}

func TestPopCount32(t *testing.T) {
	assert.Equal(t, PopCount32(0), 0)
	assert.Equal(t, PopCount32(^uint32(0)), 32)
}
