package async

func ToAsync(callback func()) chan error {
	c := make(chan error, 1)
	go func() {
		callback()
		c <- nil
	}()

	return c // <-c  receive from c to await it
}
