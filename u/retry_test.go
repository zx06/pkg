package u

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRetry(t *testing.T) {
	err := Retry(func() error {
		return fmt.Errorf("error")
	}, 3)
	if assert.Error(t, err) {
		assert.Equal(t, "error", err.Error())
	}

	err = Retry(func() error {
		return nil
	}, 3)
	assert.NoError(t, err)

	i := 0
	err = Retry(func() error {
		i++
		if i < 3 {
			return fmt.Errorf("error-%d", i)
		}
		return nil
	}, 2)
	if assert.Error(t, err) {
		assert.Equal(t, "error-2", err.Error())
		assert.Equal(t, 2, i)
	}

	i = 0
	err = Retry(func() error {
		i++
		if i < 3 {
			return fmt.Errorf("error-%d", i)
		}
		return nil
	}, 3)
	if assert.NoError(t, err) {
		assert.Equal(t, 3, i)
	}

}

func ExampleRetry() {
	i := 0
	err := Retry(func() error {
		i++
		if i < 3 {
			return fmt.Errorf("error-%d", i)
		}
		return nil
	}, 2)
	if err != nil {
		fmt.Println(err.Error())
	}
	// Output:
	// error-2
}
