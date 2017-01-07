package metadata

import (
	"sync"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

type Metadata struct {
	headers map[string]string

	mu sync.Mutex
}

const MetadataKey = "appscode-grpc-metadata"

func NewMetadata() *Metadata {
	return &Metadata{
		headers: make(map[string]string),
	}
}

func NewMetadataFromContext(ctx context.Context) *Metadata {
	meta, ok := ctx.Value(MetadataKey).(*Metadata)
	if ok {
		return meta
	}
	meta = &Metadata{
		headers: make(map[string]string),
	}
	ctx = context.WithValue(ctx, MetadataKey, meta)
	return meta
}

func (m *Metadata) AddHeader(key, value string) *Metadata {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.headers[key] = value
	return m
}

func (m *Metadata) AddHeaders(h map[string]string) *Metadata {
	m.mu.Lock()
	defer m.mu.Unlock()
	for k, v := range h {
		m.headers[k] = v
	}
	return m
}

func (m *Metadata) GRPCHeader() metadata.MD {
	m.mu.Lock()
	defer m.mu.Unlock()
	return metadata.New(m.headers)
}

func UpdateContext(ctx context.Context, m *Metadata) context.Context {
	ctx = context.WithValue(ctx, MetadataKey, m)
	return ctx
}
