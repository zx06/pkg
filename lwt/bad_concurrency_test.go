package lwt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// func TestBadLoopDataRace(t *testing.T) {
// 	r := BadLoopDataRace()
// 	assert.NotEqual(t, []int{0, 1, 2, 3, 4}, r)
// }

func TestBadLoopDataRaceFix1(t *testing.T) {
	r := BadLoopDataRaceFix1()
	assert.ElementsMatch(t, []int{0, 1, 2, 3, 4}, r)
}

func TestBadLoopDataRaceFix2(t *testing.T) {
	r := BadLoopDataRaceFix2()
	assert.ElementsMatch(t, []int{0, 1, 2, 3, 4}, r)
}
