package u

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceSplit(t *testing.T) {
	t.Run("SliceSplit 10-3", func(t *testing.T) {
		src := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		dst := SliceSplit(src, 3)
		if assert.Equal(t, 4, len(dst)) {
			assert.Equal(t, []int{1, 2, 3}, dst[0])
			assert.Equal(t, []int{4, 5, 6}, dst[1])
			assert.Equal(t, []int{7, 8, 9}, dst[2])
			assert.Equal(t, []int{10}, dst[3])
		}
	})
	t.Run("SliceSplit 10-2", func(t *testing.T) {
		src := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		dst := SliceSplit(src, 2)
		if assert.Equal(t, 5, len(dst)) {
			assert.Equal(t, []int{1, 2}, dst[0])
			assert.Equal(t, []int{3, 4}, dst[1])
			assert.Equal(t, []int{5, 6}, dst[2])
			assert.Equal(t, []int{7, 8}, dst[3])
			assert.Equal(t, []int{9, 10}, dst[4])
		}
	})
	t.Run("SliceSplit 5-10", func(t *testing.T) {
		src := []int{1, 2, 3, 4, 5}
		dst := SliceSplit(src, 10)
		if assert.Equal(t, 1, len(dst)) {
			assert.Equal(t, []int{1, 2, 3, 4, 5}, dst[0])
		}
	})
}
