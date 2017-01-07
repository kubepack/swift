package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestHeader(t *testing.T) {
	m := NewMetadata()
	m.AddHeader("foo", "bar")

	ctx := context.Background()
	ctx = UpdateContext(ctx, m)

	meta := NewMetadataFromContext(ctx)
	assert.Equal(t, meta.headers, m.headers)

	meta.AddHeader("new", "updated")
	assert.Equal(t, meta.headers, m.headers)
}
