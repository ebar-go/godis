package godis

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewServer(t *testing.T) {
	srv := NewServer()
	err := srv.Set("foo", "bar")
	if err != nil {
		t.Fatal(err)
	}

	res, err := srv.Get("foo")
	assert.Equal(t, "bar", res)

	assert.True(t, srv.Exists("foo"))
}
