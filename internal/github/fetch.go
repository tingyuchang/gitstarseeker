package github

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/spf13/viper"
)

func FetchBatchRepositories(ctx context.Context) []Repository {
	var result []Repository
	// total request make
	var numOfRequests int
	// timeout for fetch all repositories
	var timeout time.Duration

	repos := ReadSource()
	defaultTimeOut := time.Duration(int64(viper.Get("http.reqtimeout").(int))) * time.Millisecond
	timeout = time.Duration(int64(len(repos))) * defaultTimeOut
	// WARNING: setup Github auth token, or get error
	numOfRequests = len(repos)
	repoChan := make(chan int)
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	mu := sync.RWMutex{}
	wg := sync.WaitGroup{}

	// create worker
	workerCount := viper.Get("http.workernums").(int)
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range repoChan {
				ctx, cancel := context.WithTimeout(ctx, defaultTimeOut)

				repo, err := GetRepo(ctx, repos[v])
				if err != nil {
					fmt.Printf("failed: %v\terror: %v\n", repos[v], err)
					cancel()
					continue
				}
				log.Println("Get", repo.FullName)
				mu.Lock()
				result = append(result, repo)
				mu.Unlock()
				cancel()
			}
		}()
	}

	go func() {
		for i := 0; i < numOfRequests; i++ {
			repoChan <- i
		}
		close(repoChan)
	}()

	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("Fetch timeout end")
		}
	}()

	wg.Wait()

	return result
}
