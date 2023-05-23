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

	zip, err := r.DownloadVersionZip(context.Background(), "github.com/aliyun/aliyun-oss-go-sdk", "v2.2.4+incompatible")
	if err != nil {
		panic(err)
	}

	fmt.Println(len(zip))

	// Output:
	// 861606

}
