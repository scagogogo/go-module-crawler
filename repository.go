package go_module_crawler

import (
	"fmt"
	"net/url"
	"strings"
)

// ------------------------------------------------- --------------------------------------------------------------------

const (
	DefaultIndexServerURL = "https://index.golang.org"

	DefaultProxyServerURL = "https://proxy.golang.org"
)

// ProxyServerURLGoProxyCN 在国内使用这个速度会更快
const ProxyServerURLGoProxyCN = "https://goproxy.cn"

// ------------------------------------------------- --------------------------------------------------------------------

type RepositoryOptions struct {
	IndexServerURL string
	ProxyServerURL string
	ProxyIP        string
}

func NewRepositoryOptionsWithGoProxyCN() *RepositoryOptions {
	return NewRepositoryOptions().WithProxyServerURL(ProxyServerURLGoProxyCN)
}

func NewRepositoryOptions() *RepositoryOptions {
	return &RepositoryOptions{
		IndexServerURL: DefaultIndexServerURL,
		ProxyServerURL: DefaultProxyServerURL,
	}
}

func (x *RepositoryOptions) WithIndexServerURL(indexServerURL string) *RepositoryOptions {
	x.IndexServerURL = indexServerURL
	return x
}

func (x *RepositoryOptions) WithProxyServerURL(proxyServerURL string) *RepositoryOptions {
	x.ProxyServerURL = proxyServerURL
	return x
}

func (x *RepositoryOptions) WithProxyIP(proxyIP string) *RepositoryOptions {
	x.ProxyIP = proxyIP
	return x
}

// ------------------------------------------------- --------------------------------------------------------------------

// Repository 表示一个golang包的仓库
type Repository struct {
	options *RepositoryOptions
}

// NewRepository 创建一个仓库，如果不传递ServerURL的话则默认使用官方的golang.org
func NewRepository(options ...*RepositoryOptions) (*Repository, error) {

	if len(options) == 0 {
		// 没有设置ServerURL的话使用默认的
		options = append(options, NewRepositoryOptions())
	} else {
		// 设置了的话格式化统一，同时进行参数检查
		options[0].ProxyServerURL = formatServerURL(options[0].ProxyServerURL)
		_, err := url.Parse(options[0].ProxyServerURL)
		if err != nil {
			return nil, fmt.Errorf("proxy server url %s error: %s", options[0].ProxyServerURL, err.Error())
		}

		options[0].IndexServerURL = formatServerURL(options[0].IndexServerURL)
		_, err = url.Parse(options[0].IndexServerURL)
		if err != nil {
			return nil, fmt.Errorf("index server url %s error: %s", options[0].IndexServerURL, err.Error())
		}

	}

	return &Repository{
		options: options[0],
	}, nil
}

// 对仓库URL进行格式统一，方便后面使用
func formatServerURL(serverUrl string) string {
	return strings.TrimRight(serverUrl, "/")
}
