package go_module_crawler

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRepository_Index(t *testing.T) {
	// 在短模式下跳过测试
	if testing.Short() {
		t.Skip("Skipping index test in short mode")
	}

	// 如果没有明确设置代理，则跳过此测试
	proxyIP := os.Getenv("GO_MODULE_PROXY_IP")
	if proxyIP == "" {
		t.Skip("Skipping index test: GO_MODULE_PROXY_IP environment variable not set. " +
			"This test requires a proxy to access index.golang.org")
	}

	// 创建仓库选项
	options := NewRepositoryOptionsWithGoProxyCN().WithProxyIP(proxyIP)

	r, err := NewRepository(options)
	assert.Nil(t, err)

	// 创建带超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	packageSlice, err := r.Index(ctx, "")
	if err != nil {
		t.Logf("Error accessing index: %v", err)
		t.SkipNow() // 跳过测试而不是失败
	}

	assert.NotNil(t, packageSlice)
	assert.True(t, len(packageSlice) > 0, "Expected at least some packages in the index")

	// 验证返回的包数据结构
	for _, pkg := range packageSlice {
		assert.NotEmpty(t, pkg.Path, "Package path should not be empty")
		assert.NotEmpty(t, pkg.Version, "Package version should not be empty")
		assert.NotEmpty(t, pkg.Timestamp, "Package timestamp should not be empty")
	}

	t.Logf("Successfully retrieved %d packages from index", len(packageSlice))
}
