package go_module_crawler

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepository_GetLatestVersion(t *testing.T) {
	r, err := NewRepository(NewRepositoryOptionsWithGoProxyCN())
	assert.Nil(t, err)

	version, err := r.GetLatestVersionInformation(context.Background(), "github.com/aliyun/aliyun-oss-go-sdk")
	assert.Nil(t, err)
	t.Log(version)

}

func TestRepository_ListVersions(t *testing.T) {
	r, err := NewRepository(NewRepositoryOptionsWithGoProxyCN())
	assert.Nil(t, err)

	versions, err := r.ListVersions(context.Background(), "k8s.io/apimachinery")
	assert.Nil(t, err)
	assert.True(t, len(versions) > 100)
	t.Log(versions)
}

func TestRepository_GetVersionInformation(t *testing.T) {
	r, err := NewRepository(NewRepositoryOptionsWithGoProxyCN())
	assert.Nil(t, err)

	version, err := r.GetVersionInformation(context.Background(), "k8s.io/apimachinery", "v0.26.3-rc.0")
	assert.Nil(t, err)
	t.Log(version)
}
