# Go Module Crawler

# 一、这是什么？

`Go Module`爬虫，用于获取`Go`的模块信息。

# 二、安装依赖

```bash
go get -u github.com/scagogogo/go-module-crawler
```

# 三、API 示例

## 3.1 创建仓库实例 

首先需要介绍一下仓库的概念，我们认为仓库是一个托管`go module`的地方，通常情况下指的是`https://proxy.golang.org`和`https://index.golang.org`，仓库提供了两个功能：

- 通过`Proxy Server URL`来获取模块和它的版本的相关信息或者安装包
- 通过`Index Server URL`根据时间排序来遍历整个仓库中的所有的包 

仓库需要设置两个必须的参数：

- `Proxy Server URL` ，默认值是`https://proxy.golang.org`，国内用户推荐设置为`https://goproxy.cn`
- `Index Server UR`L ，默认值是`https://index.golang.org`

还有一个可选的参数：

- `ProxyIP`，这是因为`proxy.golang.org`和`index.golang.org`都需要翻墙才能访问，这个参数是用来在国内挂代理访问的。

```go
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
```

## 3.2 获取模块的最新版本

根据模块名称获取此模块的最新版本： 

```go
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

```

## 3.3 列出模块的所有版本

根据模块名称列出模块的所有版本：

```go
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
```

## 3.4 列出模块的某个版本的信息

根据模块名称和版本列出此版本的相关信息： 

```go
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
```

## 3.5 获取模块的指定版本的go.mod文件

根据模块名称和版本获取这个版本的`go.mod`文件的内容： 

```go
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

	mod, err := r.GetGoMod(context.Background(), "k8s.io/apimachinery", "v0.26.3-rc.0")
	if err != nil {
		panic(err)
	}
	fmt.Println(mod)

	// Output:
	// // This is a generated file. Do not edit directly.
	//
	// module k8s.io/apimachinery
	//
	// go 1.19
	//
	// require (
	//        github.com/armon/go-socks5 v0.0.0-20160902184237-e75332964ef5
	//        github.com/davecgh/go-spew v1.1.1
	//        github.com/elazarl/goproxy v0.0.0-20180725130230-947c36da3153
	//        github.com/evanphx/json-patch v4.12.0+incompatible
	//        github.com/gogo/protobuf v1.3.2
	//        github.com/golang/protobuf v1.5.2
	//        github.com/google/gnostic v0.5.7-v3refs
	//        github.com/google/go-cmp v0.5.9
	//        github.com/google/gofuzz v1.1.0
	//        github.com/google/uuid v1.1.2
	//        github.com/moby/spdystream v0.2.0
	//        github.com/mxk/go-flowrate v0.0.0-20140419014527-cca7078d478f
	//        github.com/spf13/pflag v1.0.5
	//        github.com/stretchr/testify v1.8.0
	//        golang.org/x/net v0.7.0
	//        gopkg.in/inf.v0 v0.9.1
	//        k8s.io/klog/v2 v2.80.1
	//        k8s.io/kube-openapi v0.0.0-20221012153701-172d655c2280
	//        k8s.io/utils v0.0.0-20221107191617-1a15be271d1d
	//        sigs.k8s.io/json v0.0.0-20220713155537-f223a00ba0e2
	//        sigs.k8s.io/structured-merge-diff/v4 v4.2.3
	//        sigs.k8s.io/yaml v1.3.0
	//)
	//
	// require (
	//        github.com/go-logr/logr v1.2.3 // indirect
	//        github.com/json-iterator/go v1.1.12 // indirect
	//        github.com/kr/text v0.2.0 // indirect
	//        github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	//        github.com/modern-go/reflect2 v1.0.2 // indirect
	//        github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	//        github.com/onsi/ginkgo/v2 v2.4.0 // indirect
	//        github.com/onsi/gomega v1.23.0 // indirect
	//        github.com/pkg/errors v0.9.1 // indirect
	//        github.com/pmezard/go-difflib v1.0.0 // indirect
	//        golang.org/x/text v0.7.0 // indirect
	//        google.golang.org/protobuf v1.28.1 // indirect
	//        gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	//        gopkg.in/yaml.v2 v2.4.0 // indirect
	//        gopkg.in/yaml.v3 v3.0.1 // indirect
	//)

}
```

## 3.6 获取模块的某个版本的安装文件

根据模块名称和版本获取这个版本的安装文件：

```go
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
```

## 3.7 遍历索引

根据索引遍历整个仓库中的所有包，包可能会比较多达到上千万，请谨慎遍历避免对服务器造成不必要的压力。

注意部分仓库不支持遍历，所以索引服务器通常使用`https://index.golang.org/`。

```go
package main

import (
	"context"
	"fmt"
	go_module_crawler "github.com/scagogogo/go-module-crawler"
	"time"
)

func main() {

	options := go_module_crawler.NewRepositoryOptions().WithProxyIP("http://127.0.0.1:7890")
	repository, err := go_module_crawler.NewRepository(options)
	if err != nil {
		panic(err)
	}

	since := ""
	for {
		packageSlice, err := repository.Index(context.Background(), since)
		if err != nil {
			panic(err)
		}
		for _, v := range packageSlice {
			fmt.Println(fmt.Sprintf("%s@%s", v.Path, v.Version))
			since = v.Timestamp
		}
		time.Sleep(time.Second)
	}

	// Output:
	// golang.org/x/text@v0.3.0
	// golang.org/x/crypto@v0.0.0-20190404164418-38d8ce5564a5
	// github.com/FiloSottile/mkcert@v1.3.0
	// github.com/DHowett/go-plist@v0.0.0-20180609054337-500bd5b9081b
	// software.sslmate.com/src/go-pkcs12@v0.0.0-20180114231543-2291e8f0f237
	// golang.org/x/net@v0.0.0-20180627171509-e514e69ffb8b
	// golang.org/x/exp/notary@v0.0.0-20190409044807-56b785ea58b2
	// golang.org/x/crypto@v0.0.0-20181025213731-e84da0312774
	// golang.org/x/net@v0.0.0-20181213202711-891ebc4b82d6
	// golang.org/x/sys@v0.0.0-20190306220234-b354f8bf4d9e
	// golang.org/x/text@v0.1.1-0.20171102144821-8253218a5ec6
	// golang.org/x/net@v0.0.0-20180906233101-161cd47e91fd
	// git.apache.org/thrift.git@v0.0.0-20180807212849-6e67faa92827
	// git.apache.org/thrift.git@v0.0.0-20180902110319-2566ecd5d999
	// github.com/beorn7/perks@v0.0.0-20180321164747-3a771d992973
	// github.com/ghodss/yaml@v1.0.0
	// github.com/golang/glog@v0.0.0-20160126235308-23def4e6c14b
	// gocloud.dev@v0.10.0
	// github.com/google/go-containerregistry@v0.0.0-20190327192230-5296537b6d5d
	// rsc.io/quote@v1.0.0
	// gocloud.dev@v0.12.0
	// ...

}

```

