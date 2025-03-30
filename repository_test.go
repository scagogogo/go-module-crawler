package go_module_crawler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRepository(t *testing.T) {
	// 测试默认选项
	repo, err := NewRepository()
	assert.Nil(t, err)
	assert.NotNil(t, repo)
	assert.Equal(t, DefaultProxyServerURL, repo.options.ProxyServerURL)
	assert.Equal(t, DefaultIndexServerURL, repo.options.IndexServerURL)
	assert.Empty(t, repo.options.ProxyIP)

	// 测试自定义选项
	customOptions := NewRepositoryOptions().
		WithProxyServerURL("https://custom-proxy.example.com").
		WithIndexServerURL("https://custom-index.example.com").
		WithProxyIP("http://proxy.example.com:8080")

	repo, err = NewRepository(customOptions)
	assert.Nil(t, err)
	assert.NotNil(t, repo)
	assert.Equal(t, "https://custom-proxy.example.com", repo.options.ProxyServerURL)
	assert.Equal(t, "https://custom-index.example.com", repo.options.IndexServerURL)
	assert.Equal(t, "http://proxy.example.com:8080", repo.options.ProxyIP)

	// 测试 GoProxyCN 选项
	proxyOptions := NewRepositoryOptionsWithGoProxyCN()
	repo, err = NewRepository(proxyOptions)
	assert.Nil(t, err)
	assert.NotNil(t, repo)
	assert.Equal(t, ProxyServerURLGoProxyCN, repo.options.ProxyServerURL)
	assert.Equal(t, DefaultIndexServerURL, repo.options.IndexServerURL)
}

func TestNewRepository_InvalidURLs(t *testing.T) {
	// 测试无效的代理服务器URL
	invalidProxyOptions := NewRepositoryOptions().WithProxyServerURL("://invalid-url")
	_, err := NewRepository(invalidProxyOptions)
	assert.NotNil(t, err, "应当返回错误，因为代理URL无效")
	assert.Contains(t, err.Error(), "proxy server url")

	// 测试无效的索引服务器URL
	invalidIndexOptions := NewRepositoryOptions().WithIndexServerURL("://invalid-url")
	_, err = NewRepository(invalidIndexOptions)
	assert.NotNil(t, err, "应当返回错误，因为索引URL无效")
	assert.Contains(t, err.Error(), "index server url")
}

func TestFormatServerURL(t *testing.T) {
	// 测试URL格式化函数
	testCases := []struct {
		input    string
		expected string
	}{
		{"https://example.com/", "https://example.com"},
		{"https://example.com", "https://example.com"},
		{"https://example.com///", "https://example.com"},
		{"", ""},
	}

	for _, tc := range testCases {
		result := formatServerURL(tc.input)
		assert.Equal(t, tc.expected, result)
	}
}
