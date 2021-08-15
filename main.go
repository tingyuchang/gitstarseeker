package main

import (
	"context"
	"fmt"
	"gitstarseeker/internal/github"
	"log"
	"time"
)

func main() {
	repos := github.ReadSource()
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	go func() {
		repo, err := github.GetRepo(ctx, repos[0])
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("Repository: %s\ndesciption: %s\nstarts: %d\n", repo.FullName, repo.Description, repo.StargazersCount)
	}()

	select {
	case <-ctx.Done():
		fmt.Println("program end.")
	}
}
