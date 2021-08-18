package main

import (
	"context"
	"fmt"
	"gitstarseeker/internal/github"
	"time"

	"github.com/spf13/viper"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()
	// setup conf
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}

	// 1. read source file
	src := github.ReadSource(viper.Get("source.awesome-go").(string))
	// 2. create ctx for batch http get
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// exec batch
	repos := github.FetchBatchRepositories(ctx, src)

	for _, v := range repos {
		fmt.Printf("Repository: %s\tdesciption: %s\tstarts: %d\n", v.FullName, v.Description, v.StargazersCount)
	}
}
