package go_module_crawler

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepository_Index(t *testing.T) {

	r, err := NewRepository(NewRepositoryOptionsWithGoProxyCN().WithProxyIP("http://127.0.0.1:7890"))
	assert.Nil(t, err)

	packageSlice, err := r.Index(context.Background(), "")
	assert.Nil(t, err)
	t.Log(packageSlice)
}
