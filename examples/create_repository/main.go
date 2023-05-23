package main

import (
	"fmt"
	go_module_crawler "github.com/scagogogo/go-module-crawler"
)

func main() {

	// 使用默认配置创建仓库
	repository, err := go_module_crawler.NewRepository()
	if err != nil {
		panic(err)
	}
	fmt.Println(repository)

	// 配置代理IP
	options := go_module_crawler.NewRepositoryOptions().WithProxyIP("http://127.0.0.1:7890")
	repository, err = go_module_crawler.NewRepository(options)
	if err != nil {
		panic(err)
	}
	fmt.Println(repository)

	// 指定proxy server 、 index server 等
	options = go_module_crawler.NewRepositoryOptions().
		WithProxyIP("http://127.0.0.1:7890").
		WithProxyServerURL(go_module_crawler.ProxyServerURLGoProxyCN)
	repository, err = go_module_crawler.NewRepository(options)
	if err != nil {
		panic(err)
	}
	fmt.Println(repository)
}
