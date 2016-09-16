package redis

import (
	"github.com/nu7hatch/gouuid"
	"testing"
)

func TestSetAndGet(t *testing.T) {
	uuid, err := uuid.NewV4()
	if err != nil {
		t.Error("Cannot generate new key")
	}
	key, value := uuid.String(), "testvalue"
	Set("test", key, value)
	v := Get("test", key)
	Del("test", key)
	if value != v {
		t.Errorf("[%s] should be %s, but it's %s.", key, value, v)
	}
}
