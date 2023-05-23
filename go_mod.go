package go_module_crawler

import (
	"context"
	"fmt"
)

// 获取go.mod文件：
// https://goproxy.cn/github.com/aliyun/aliyun-oss-go-sdk/@v/v2.2.4+incompatible.mod
// https://goproxy.cn/github.com/aliyun/aliyun-oss-go-sdk/@v/v2.2.4+incompatible.info

func (x *Repository) GetGoMod(ctx context.Context, moduleName, version string) (string, error) {
	targetUrl := x.BuildGetGoModURL(moduleName, version)
	responseBytes, err := x.Request(ctx, targetUrl)
	if err != nil {
		return "", err
	}
	return string(responseBytes), nil
}

func (x *Repository) BuildGetGoModURL(moduleName, version string) string {
	return fmt.Sprintf("%s/%s/@v/%s.mod", x.options.ProxyServerURL, moduleName, version)
}
