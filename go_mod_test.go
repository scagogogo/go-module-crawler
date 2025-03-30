package go_module_crawler

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepository_GetGoMod(t *testing.T) {
	// 在短模式下跳过网络依赖测试
	if testing.Short() {
		t.Skip("Skipping network-dependent test in short mode")
	}

	// 从环境变量获取代理地址
	proxyIP := os.Getenv("GO_MODULE_PROXY_IP")

	// 创建仓库选项
	options := NewRepositoryOptionsWithGoProxyCN()
	if proxyIP != "" {
		options = options.WithProxyIP(proxyIP)
	}

	r, err := NewRepository(options)
	assert.Nil(t, err)

	mod, err := r.GetGoMod(context.Background(), "k8s.io/apimachinery", "v0.26.3-rc.0")
	if err != nil {
		t.Logf("Error: %v. This test requires internet access to goproxy.cn", err)
		t.SkipNow()
	}

	assert.NotEmpty(t, mod, "go.mod content should not be empty")
	assert.True(t, strings.Contains(mod, "module k8s.io/apimachinery"), "go.mod should contain module definition")
	assert.True(t, strings.Contains(mod, "go "), "go.mod should contain go version")

	t.Logf("Successfully retrieved go.mod with length: %d bytes", len(mod))
}
