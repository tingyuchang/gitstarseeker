package github

import (
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func ReadSource() []string {
	c := make(chan struct{})
	var result []string
	go func() {
		// TODO: user could set any source
		res, err := http.Get("https://raw.githubusercontent.com/avelino/awesome-go/master/README.md")
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

func findGithubRepos(srcData []byte) ([]string, error) {
	// `github.com/\S+\) `
	re := regexp.MustCompile(`github.com/\S+\) `)
	found := re.FindAll(srcData, -1)
	var returnData []string

	for _, v := range found {

		if strings.Contains(string(v), "CONTRIBUTING.md") {
			continue
		}
		// escape "github.com/" and ") "
		returnData = append(returnData, string(v[11:len(v)-2]))
	}
	return returnData, nil
}
