package go_module_crawler

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepository_GetGoMod(t *testing.T) {
	r, err := NewRepository(NewRepositoryOptionsWithGoProxyCN())
	assert.Nil(t, err)

	mod, err := r.GetGoMod(context.Background(), "k8s.io/apimachinery", "v0.26.3-rc.0")
	assert.Nil(t, err)
	assert.NotEmpty(t, mod)
}
