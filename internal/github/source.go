package github

import (
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
)
// ReadSource read source from url
// return []string which store path of github repository
func ReadSource(url string) []string {
	c := make(chan struct{})
	var result []string
	go func() {
		// TODO: user could set any type source ex: http or file
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		data, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		repos, err := findGithubRepos(data)
		result = repos
		c <- struct{}{}
	}()

	select {
	case <-c:
		return result
	}
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
