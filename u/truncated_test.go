package u

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTruncatedString(t *testing.T) {
	t.Run("src is empty", func(t *testing.T) {
		r := TruncatedString("", 10)
		assert.Equal(t, "", r)
	})
	t.Run("src is less than limit", func(t *testing.T) {
		r := TruncatedString("abc", 10)
		assert.Equal(t, "abc", r)
	})
	t.Run("src is more than limit", func(t *testing.T) {
		r := TruncatedString("abcdefghijklmnopqrstuvwxyz", 10)
		assert.Equal(t, "abcdefg...", r)
	})
	t.Run("src is equal to limit", func(t *testing.T) {
		r := TruncatedString("abcdefghijklmnopqrstuvwxyz", 26)
		assert.Equal(t, "abcdefghijklmnopqrstuvwxyz", r)
	})
}
