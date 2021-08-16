package github

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const GITHUB_API_URL = "https://api.github.com/"

/*
For API requests using personal token,
you can make up to 5,000 requests per hour.
For unauthenticated requests, the rate limit allows for up to 60 requests per hour.
*/

func httpDo(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error {
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

func GetRepo(ctx context.Context, userRepo string) (Repository, error) {
	// setup request
	req, err := http.NewRequest("GET", GITHUB_API_URL+"repos/"+userRepo, nil)
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("Authorization", "token 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")
	if err != nil {
		return Repository{}, err
	}
	repo := Repository{}
	err = httpDo(ctx, req, func(resp *http.Response, err error) error {
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		err = json.Unmarshal(body, &repo)
		if err != nil {
			return err
		}
		return nil
	})
	return repo, err
}
