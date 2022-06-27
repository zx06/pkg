package u


// Retry calls the given function until it returns nil or the given number of retries is reached.
func Retry(f func() error, maxRetries int) error {
	var err error
	for i := 0; i < maxRetries; i++ {
		err = f()
		if err == nil {
			return nil
		}
	}
	return err
}
