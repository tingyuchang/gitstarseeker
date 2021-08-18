package github

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/spf13/viper"
)
// FetchBatchRepositories loads git repository data (ex: mewkiz/flac) from src
// organize a batch requests group using concept about simple worker (no pool)
// worker handle request from repoChan, it's chan int, worker use this int value to find repo address from src
// each request has timeout context for cancellation
func FetchBatchRepositories(ctx context.Context, src []string) []Repository {
	var result []Repository
	// total request make
	var numOfRequests int
	// timeout for fetch all repositories
	var timeout time.Duration

	defaultTimeOut := time.Duration(int64(viper.Get("http.reqtimeout").(int))) * time.Millisecond
	timeout = time.Duration(int64(len(src))) * defaultTimeOut
	// WARNING: setup Github auth token, or get error
	numOfRequests = len(src)
	numOfRequests = 20
	repoChan := make(chan int)
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	mu := sync.RWMutex{}
	wg := sync.WaitGroup{}

	// create worker, get worker count form conf
	workerCount := viper.Get("http.workernums").(int)
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range repoChan {
				ctx, cancel := context.WithTimeout(ctx, defaultTimeOut)

				repo, err := GetRepo(ctx, src[v])
				if err != nil {
					fmt.Printf("failed: %v\terror: %v\n", src[v], err)
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
