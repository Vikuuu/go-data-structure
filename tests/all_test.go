package tests

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"

	ds "github.com/Vikuuu/go-datastructure"
)

func TestBinarySearch(t *testing.T) {}

func TestTwoCrystalBalls(t *testing.T) {
	idx := rand.Intn(10000)
	data := make([]bool, 10000)

	assert.Equal(t, ds.TwoCrystalBalls(data), -1)
	for i := idx; i < 10000; i++ {
		data[i] = true
	}

	assert.Equal(t, ds.TwoCrystalBalls(data), idx)
}
