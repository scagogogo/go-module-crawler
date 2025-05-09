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
