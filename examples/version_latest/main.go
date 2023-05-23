package main

import (
	"context"
	"fmt"
	go_module_crawler "github.com/scagogogo/go-module-crawler"
)

func main() {

	r, err := go_module_crawler.NewRepository(go_module_crawler.NewRepositoryOptionsWithGoProxyCN())
	if err != nil {
		panic(err)
	}

	version, err := r.GetLatestVersionInformation(context.Background(), "github.com/aliyun/aliyun-oss-go-sdk")
	if err != nil {
		panic(err)
	}

	fmt.Println(version.Version)
	fmt.Println(version.Time)

	// Output:
	// v2.2.7+incompatible
	// 2023-03-23 08:33:02 +0000 UTC
}
