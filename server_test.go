package godis

import (
	"fmt"
	"github.com/ebar-go/godis/constant"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewServer(t *testing.T) {
	srv := NewServer()
	key := "foo"
	srv.Set(key, "bar")

	assert.Equal(t, "bar", srv.Get(key))

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

func TestHExists(t *testing.T) {
	srv := NewServer()
	srv.HSet("someHash", "foo", "bar")

	assert.True(t, srv.HExists("someHash", "foo"))
	assert.False(t, srv.HExists("someHash", "notExist"))
}

func TestHLen(t *testing.T) {
	srv := NewServer()
	srv.HSet("someHash", "foo", "bar")
	assert.Equal(t, int64(1), srv.HLen("someHash"))
}

func TestHDel(t *testing.T) {
	srv := NewServer()
	srv.HSet("someHash", "foo", "bar")

	srv.HSet("someHash", "age", 1)
	assert.Equal(t, 0, srv.HDel("someHash", "notExistField"))
	assert.Equal(t, 2, srv.HDel("someHash", "foo", "age"))
	assert.Equal(t, int64(0), srv.HLen("someHash"))
}

func TestHKeys(t *testing.T) {
	srv := NewServer()
	srv.HSet("someHash", "foo", "bar")

	srv.HSet("someHash", "age", 1)
	assert.NotEmpty(t, srv.HKeys("someHash"))
	assert.Nil(t, srv.HKeys("notExistKey"))
}

func TestHGetAll(t *testing.T) {
	srv := NewServer()
	srv.HSet("someHash", "foo", "bar")

	srv.HSet("someHash", "age", 1)
	assert.NotEmpty(t, srv.HGetAll("someHash"))
	assert.Nil(t, srv.HGetAll("notExistKey"))
}

func TestSAdd(t *testing.T) {
	srv := NewServer()
	key := "someSet"
	assert.Nil(t, srv.SAdd(key, "foo", "bar"))

	strKey := "age"
	srv.Set(strKey, 1)
	assert.NotNil(t, srv.SAdd(strKey, "foo", "bar"))
}

func TestSRem(t *testing.T) {
	srv := NewServer()
	key := "someSet"
	assert.Nil(t, srv.SAdd(key, "foo", "bar"))
	count, err := srv.SRem(key, "foo", "age")
	assert.Nil(t, err)
	assert.Equal(t, 1, count)

	strKey := "age"
	srv.Set(strKey, 1)

	count, err = srv.SRem(strKey, "foo", "bar")
	assert.Equal(t, 0, count)
	assert.NotNil(t, err)
}

func TestSCard(t *testing.T) {
	srv := NewServer()
	key := "someSet"
	srv.SAdd(key, "foo", "bar")
	assert.Equal(t, int64(2), srv.SCard(key))
}

func TestSPop(t *testing.T) {
	srv := NewServer()
	key := "someSet"
	srv.SAdd(key, "foo", "bar")
	items := srv.SPop(key, 1)
	assert.NotEmpty(t, items)
	fmt.Println(items)
	assert.Equal(t, int64(1), srv.SCard(key))
}

func TestSIsMember(t *testing.T) {
	srv := NewServer()
	key := "someSet"
	srv.SAdd(key, "foo", "bar")
	assert.Equal(t, 1, srv.SIsMember(key, "foo"))
	assert.Equal(t, 0, srv.SIsMember(key, "notExistMember"))
	assert.Equal(t, 0, srv.SIsMember("notExistKey", "foo"))
}

func TestSMembers(t *testing.T) {
	srv := NewServer()
	key := "someSet"
	srv.SAdd(key, "foo", "bar")
	assert.Equal(t, 2, len(srv.SMembers(key)))
}

func TestLPush(t *testing.T) {
	srv := NewServer()
	key := "someList"
	assert.Equal(t, 2, srv.LPush(key, "foo", "bar"))

}

func TestRPush(t *testing.T) {
	srv := NewServer()
	key := "someList"
	assert.Equal(t, 2, srv.LPush(key, "foo", "bar"))
}

func TestLLen(t *testing.T) {
	srv := NewServer()
	key := "someList"

	srv.LPush(key, "foo", "bar")
	assert.Equal(t, uint64(2), srv.LLen(key))

	srv.RPush(key, "foo2")
	assert.Equal(t, uint64(3), srv.LLen(key))
}

func TestLPop(t *testing.T) {
	srv := NewServer()
	key := "someList"

	srv.LPush(key, "foo", "bar")
	fmt.Println(srv.LPop(key, 1))
	assert.Equal(t, uint64(1), srv.LLen(key))
}

func TestRPop(t *testing.T) {
	srv := NewServer()
	key := "someList"

	srv.LPush(key, "foo", "bar")
	fmt.Println(srv.RPop(key, 1))
	assert.Equal(t, uint64(1), srv.LLen(key))
}

func TestLRange(t *testing.T) {
	srv := NewServer()
	key := "someList"

	srv.LPush(key, "foo", "bar")
	fmt.Println(srv.LRange(key, 0, 1))
}

func TestZAdd(t *testing.T) {
	srv := NewServer()
	key := "someList"

	assert.Equal(t, 1, srv.ZAdd(key, "foo", 123))
	assert.Equal(t, 1, srv.ZAdd(key, "bar", 456))
}

func TestZCard(t *testing.T) {
	srv := NewServer()
	key := "someList"

	srv.ZAdd(key, "foo", 123)
	assert.Equal(t, int64(1), srv.ZCard(key))

}

func TestZRem(t *testing.T) {
	srv := NewServer()
	key := "someList"

	srv.ZAdd(key, "foo", 123)
	srv.ZAdd(key, "bar", 456)
	assert.Equal(t, 1, srv.ZRem(key, "foo"))
}
