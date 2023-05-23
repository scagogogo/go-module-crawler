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

	version, err := r.GetVersionInformation(context.Background(), "k8s.io/apimachinery", "v0.26.3-rc.0")
	if err != nil {
		panic(err)
	}
	fmt.Println(version.Version)
	fmt.Println(version.Time)

	// Output:
	// v0.26.3-rc.0
	// 2023-02-15 10:21:21 +0000 UTC

}
