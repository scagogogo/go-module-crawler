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
