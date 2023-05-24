package go_module_crawler

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// ------------------------------------------------- --------------------------------------------------------------------

// 列出包的最新版本

// https://goproxy.cn/github.com/aliyun/aliyun-oss-go-sdk/@latest

// ModuleVersionInformation 表示一个Golang模块的最新的版本
type ModuleVersionInformation struct {

	// 最新版本的版本号是多少，比如： v2.2.7+incompatible
	Version string `json:"Version"`

	// 这个最新版本的发布时间是啥时候，比如：2023-03-23T08:33:02Z
	Time time.Time `json:"Time"`

	Origin *Origin `json:"Origin"`
}

type Origin struct {
	VCS  string `json:"VCS"`
	URL  string `json:"URL"`
	Ref  string `json:"Ref"`
	Hash string `json:"Hash"`
}

// GetLatestVersionInformation 获取给定模块的最新版本
func (x *Repository) GetLatestVersionInformation(ctx context.Context, moduleName string) (*ModuleVersionInformation, error) {
	// {
	//    "Package": "v2.2.7+incompatible",
	//    "Time": "2023-03-23T08:33:02Z",
	//    "Origin": {
	//        "VCS": "git",
	//        "URL": "https://github.com/aliyun/aliyun-oss-go-sdk",
	//        "Ref": "refs/tags/v2.2.7",
	//        "Hash": "77977ff44f387fff867be985b984251ee8012529"
	//    }
	// }
	targetUrl := x.BuildGetLatestVersionURL(moduleName)
	responseBytes, err := x.Request(ctx, targetUrl)
	if err != nil {
		return nil, err
	}
	r := &ModuleVersionInformation{}
	err = json.Unmarshal(responseBytes, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (x *Repository) BuildGetLatestVersionURL(moduleName string) string {
	return fmt.Sprintf("%s/%s/@latest", x.options.ProxyServerURL, moduleName)
}

// ------------------------------------------------- --------------------------------------------------------------------

// ListVersions 按顺序列出给定的包的所有版本
// https://proxy.golang.org/k8s.io/apimachinery/@v/list
// https://goproxy.cn/{包名}/@v/list
func (x *Repository) ListVersions(ctx context.Context, moduleName string) ([]string, error) {
	targetUrl := x.BuildListVersionsURL(moduleName)
	responseBytes, err := x.Request(ctx, targetUrl)
	if err != nil {
		return nil, err
	}
	split := strings.Split(string(responseBytes), "\n")
	versions := make([]string, 0)
	for _, v := range split {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}
		versions = append(versions, v)
	}
	return versions, nil
}

func (x *Repository) BuildListVersionsURL(moduleName string) string {
	return fmt.Sprintf("%s/%s/@v/list", x.options.ProxyServerURL, moduleName)
}

// ------------------------------------------------- --------------------------------------------------------------------

func (x *Repository) GetVersionInformation(ctx context.Context, moduleName, version string) (*ModuleVersionInformation, error) {
	targetUrl := x.BuildGetVersionInformation(moduleName, version)
	responseBytes, err := x.Request(ctx, targetUrl)
	if err != nil {
		return nil, err
	}
	r := &ModuleVersionInformation{}
	err = json.Unmarshal(responseBytes, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (x *Repository) BuildGetVersionInformation(moduleName, version string) string {
	return fmt.Sprintf("%s/%s/@v/%s.info", x.options.ProxyServerURL, moduleName, version)
}

// ------------------------------------------------- --------------------------------------------------------------------

// 下载对应版本的安装包：
// https://goproxy.cn/github.com/aliyun/aliyun-oss-go-sdk/@v/v2.2.4+incompatible.zip
//

func (x *Repository) DownloadVersionZip(ctx context.Context, moduleName, version string) ([]byte, error) {
	targetUrl := x.BuildDownloadVersionZipURL(moduleName, version)
	return x.Request(ctx, targetUrl)
}

func (x *Repository) BuildDownloadVersionZipURL(moduleName, version string) string {
	return fmt.Sprintf("%s/%s/@v/%s.zip", x.options.ProxyServerURL, moduleName, version)
}

// ------------------------------------------------- --------------------------------------------------------------------
