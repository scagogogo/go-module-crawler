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

	versions, err := r.ListVersions(context.Background(), "k8s.io/apimachinery")
	if err != nil {
		panic(err)
	}
	for _, v := range versions {
		fmt.Println(v)
	}

	// Output:
	// v0.15.7
	// v0.15.8-beta.0
	// v0.15.8-beta.1
	// v0.15.8
	// v0.15.9-beta.0
	// v0.15.9
	// v0.15.10-beta.0
	// v0.15.10
	// v0.15.11-beta.0
	// v0.15.11
	// v0.15.12-beta.0
	// v0.15.12
	// v0.15.13-beta.0
	// v0.16.4
	// v0.16.5-beta.0
	// v0.16.5-beta.1
	// v0.16.5
	// v0.16.6-beta.0
	// v0.16.6
	// ...
}
