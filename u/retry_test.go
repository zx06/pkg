package u

import (
	"fmt"
	"testing"
)

func TestRetry(t *testing.T) {
	var err error
	var i int
	var f func() error

	f = func() error {
		i++
		if i < 3 {
			return fmt.Errorf("error %d", i)
		}
		return nil
	}

	err = Retry(f, 3)
	if err != nil {
		t.Errorf("Retry returned error: %v", err)
	}
}
