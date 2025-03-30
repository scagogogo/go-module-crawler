package go_module_crawler

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// 创建带合适代理的仓库实例
func createTestRepository(t *testing.T) *Repository {
	// 从环境变量获取代理地址
	proxyIP := os.Getenv("GO_MODULE_PROXY_IP")

	// 创建仓库选项
	options := NewRepositoryOptionsWithGoProxyCN()
	if proxyIP != "" {
		options = options.WithProxyIP(proxyIP)
	}

	r, err := NewRepository(options)
	assert.Nil(t, err)
	return r
}

func TestRepository_GetLatestVersion(t *testing.T) {
	// 在短模式下跳过网络依赖测试
	if testing.Short() {
		t.Skip("Skipping network-dependent test in short mode")
	}

	r := createTestRepository(t)

	version, err := r.GetLatestVersionInformation(context.Background(), "github.com/aliyun/aliyun-oss-go-sdk")
	if err != nil {
		t.Logf("Error: %v. This test requires internet access to goproxy.cn", err)
		t.SkipNow()
	}

	assert.NotEmpty(t, version.Version, "Version should not be empty")
	assert.False(t, version.Time.IsZero(), "Time should be valid")
	t.Logf("Latest version: %s, published at: %s", version.Version, version.Time.Format(time.RFC3339))
}

func TestRepository_ListVersions(t *testing.T) {
	// 在短模式下跳过网络依赖测试
	if testing.Short() {
		t.Skip("Skipping network-dependent test in short mode")
	}

	r := createTestRepository(t)

	versions, err := r.ListVersions(context.Background(), "k8s.io/apimachinery")
	if err != nil {
		t.Logf("Error: %v. This test requires internet access to goproxy.cn", err)
		t.SkipNow()
	}

	assert.NotEmpty(t, versions, "Should return at least one version")
	assert.True(t, len(versions) > 10, "Expected multiple versions")

	// 检查版本格式是否正确
	for _, v := range versions {
		assert.NotEmpty(t, v, "Version should not be empty")
		assert.Contains(t, v, "v", "Version should contain 'v' prefix")
	}

	t.Logf("Found %d versions", len(versions))
}

func TestRepository_GetVersionInformation(t *testing.T) {
	// 在短模式下跳过网络依赖测试
	if testing.Short() {
		t.Skip("Skipping network-dependent test in short mode")
	}

	r := createTestRepository(t)

	version, err := r.GetVersionInformation(context.Background(), "k8s.io/apimachinery", "v0.26.3-rc.0")
	if err != nil {
		t.Logf("Error: %v. This test requires internet access to goproxy.cn", err)
		t.SkipNow()
	}

	assert.Equal(t, "v0.26.3-rc.0", version.Version, "Version should match requested version")
	assert.False(t, version.Time.IsZero(), "Time should be valid")
	t.Logf("Version info: %s, published at: %s", version.Version, version.Time.Format(time.RFC3339))
}

func TestRepository_DownloadVersionZip(t *testing.T) {
	// 在短模式下跳过网络依赖测试
	if testing.Short() {
		t.Skip("Skipping network-dependent test in short mode")
	}

	r := createTestRepository(t)

	zip, err := r.DownloadVersionZip(context.Background(), "github.com/aliyun/aliyun-oss-go-sdk", "v2.2.4+incompatible")
	if err != nil {
		t.Logf("Error: %v. This test requires internet access to goproxy.cn", err)
		t.SkipNow()
	}

	assert.NotNil(t, zip, "ZIP data should not be nil")
	assert.True(t, len(zip) > 1000, "ZIP data should have substantial size")
	t.Logf("Downloaded ZIP size: %d bytes", len(zip))
}
