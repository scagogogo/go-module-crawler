package go_module_crawler

import (
	"context"
	"fmt"
)

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
