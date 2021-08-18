package github

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"gitstarseeker/internal/service"
	"io/ioutil"
	"net/http"
)

const GITHUB_API_URL = "https://api.github.com/"

/*
For API requests using personal token,
you can make up to 5,000 requests per hour.
For unauthenticated requests, the rate limit allows for up to 60 requests per hour.
*/

func GetRepo(ctx context.Context, userRepo string) (Repository, error) {
	// setup request
	req, err := http.NewRequest("GET", GITHUB_API_URL+"repos/"+userRepo, nil)
	req.Header.Set("Accept", viper.Get("http.githubheaderaccept").(string))
	req.Header.Set("Authorization", fmt.Sprintf("token %v", viper.Get("http.githubheaderauthorization")))
	if err != nil {
		return Repository{}, err
	}
	repo := Repository{}
	err = service.HttpDo(ctx, req, func(resp *http.Response, err error) error {
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
