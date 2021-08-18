package service

import (
	"context"
	"net/http"
)

func HttpDo(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error {
	c := make(chan error, 1)
	defer close(c)
	req = req.WithContext(ctx)
	// run http request on goroutine, and pass receive values to f
	go func() {
		c <- f(http.DefaultClient.Do(req))
	}()

	select {
	case <-ctx.Done():
		// context is end, but need waiting f() to return
		// otherwise running f() will send to close channel
		<-c
		return ctx.Err()
	case err := <-c:
		// normal case, receive from f()
		return err
	}
}