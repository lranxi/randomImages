package cache

import (
	"testing"
	"time"
)

func TestCache_Get(t *testing.T) {
	client, err := NewCache()
	if err != nil {
		t.Fatal(err)
	}
	client.Set("test", "test", time.Second*100)
}
