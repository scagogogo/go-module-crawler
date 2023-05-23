package go_module_crawler

import (
	"context"
	"crypto/tls"
	"github.com/crawler-go-go-go/go-requests"
	"net/http"
	"net/url"
)

func (x *Repository) Request(ctx context.Context, targetUrl string) ([]byte, error) {
	options := requests.NewOptions[any, []byte](targetUrl, func(httpResponse *http.Response) ([]byte, error) {
		// TODO 增加对异常响应的拦截处理
		return requests.BytesResponseHandler(200)(httpResponse)
	})

	if x.options.ProxyIP != "" {
		proxyIP, err := url.Parse(x.options.ProxyIP)
		if err != nil {
			return nil, err
		}
		options.AppendRequestSetting(func(client *http.Client, request *http.Request) error {
			client.Transport = &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
				Proxy: http.ProxyURL(proxyIP),
			}
			return nil
		})
	}

	return requests.GetBytes(ctx, targetUrl, options)
}
