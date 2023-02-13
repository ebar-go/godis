package godis

import (
	"github.com/ebar-go/godis/constant"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewServer(t *testing.T) {
	srv := NewServer()
	key := "foo"
	err := srv.Set(key, "bar")
	if err != nil {
		t.Fatal(err)
	}

	res, err := srv.Get(key)
	assert.Equal(t, "bar", res)

	assert.True(t, srv.Exists(key))

	assert.Equal(t, constant.ExpireResultOfForever, srv.TTL(key))

	assert.Nil(t, srv.Expire(key, 10))
	assert.Equal(t, int64(10), srv.TTL(key))

	time.Sleep(time.Second * 10)
	assert.Equal(t, constant.ExpireResultOfExpired, srv.TTL(key))

	assert.Equal(t, uint(1), srv.Del(key))
	assert.False(t, srv.Exists(key))
}

func TestServer_HSet(t *testing.T) {
	srv := NewServer()
	assert.Nil(t, srv.HSet("someHash", "foo", "bar"))
	assert.Nil(t, srv.HSet("someHash", "age", 1))

	stringVal, _ := srv.HGet("someHash", "foo")
	assert.Equal(t, "bar", stringVal)

	intVal, _ := srv.HGet("someHash", "age")
	assert.Equal(t, 1, intVal)

}
