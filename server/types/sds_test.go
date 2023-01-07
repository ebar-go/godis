package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSDS(t *testing.T) {
	sds := NewSDS("foo")
	assert.NotNil(t, sds)
}

func TestSDS_Len(t *testing.T) {
	sds := NewSDS("foo")
	assert.Equal(t, uint(3), sds.Len())
}

func TestSDS_String(t *testing.T) {
	sds := NewSDS("foo")
	assert.Equal(t, "foo", sds.String())
}

func BenchmarkNewSDS(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = NewSDS("foo")
	}
}
