package main

import (
	"context"
	"fmt"
	"gitstarseeker/internal/github"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	repos := github.FetchBatchRepositories(ctx)

	for _, v := range repos {
		fmt.Printf("Repository: %s\tdesciption: %s\tstarts: %d\n", v.FullName, v.Description, v.StargazersCount)
	}
}
