package github

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const GITHUB_API_URL = "https://api.github.com/"

func GetRepo(ctx context.Context, userRepo string) (Repository, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", GITHUB_API_URL+"repos/"+userRepo, nil)
	if err != nil {
		log.Println(err)
		return Repository{}, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return Repository{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return Repository{}, err
	}

	repo := Repository{}
	err = json.Unmarshal(body, &repo)
	if err != nil {
		log.Println(err)
		return Repository{}, err
	}
	return repo, nil
}
