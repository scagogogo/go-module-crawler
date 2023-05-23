package go_module_crawler

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepository_DownloadVersionZip(t *testing.T) {
	r, err := NewRepository(NewRepositoryOptionsWithGoProxyCN())
	assert.Nil(t, err)

	zip, err := r.DownloadVersionZip(context.Background(), "github.com/aliyun/aliyun-oss-go-sdk", "v2.2.4+incompatible")
	assert.Nil(t, err)
	assert.True(t, len(zip) > 0)
}
