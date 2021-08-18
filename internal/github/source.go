package github

import (
	"context"
	"gitstarseeker/internal/service"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
)
// ReadSource read source from url
// return []string which store path of github repository
func ReadSource(url string) []string {
	// TODO: user could set any type source ex: http or file
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Print(err)
		return nil
	}
	var result []string
	err = service.HttpDo(ctx, req, func(resp *http.Response, err error) error {
		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		repos, err := findGithubRepos(data)
		if err != nil {
			return err
		}
		result = repos
		return nil
	})

	return result
}

// findGithubRepos execute regexp to fid valid github repository path
func findGithubRepos(srcData []byte) ([]string, error) {
	// `github.com/\S+\) `
	re := regexp.MustCompile(`github.com/\S+\) `)
	found := re.FindAll(srcData, -1)
	var returnData []string

	for _, v := range found {
		// remove invalid data
		// TODO: need more smart
		if strings.Contains(string(v), "CONTRIBUTING.md") {
			continue
		}
		// escape "github.com/" and ") "
		returnData = append(returnData, string(v[11:len(v)-2]))
	}
	return returnData, nil
}
