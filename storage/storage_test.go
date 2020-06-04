package storage

import (
	"fmt"
	"testing"
	"time"
)

func TestMemory1(t *testing.T) {
	Memory().Set("key1", "val1", time.Second)
	if val, err := Memory().Get("key1"); err != nil {
		t.Error()
	} else {
		fmt.Println(val)
	}
	time.Sleep(time.Second * 2)
	if val, err := Memory().Get("key1"); err != nil {
		t.Error(err)
	} else {
		fmt.Println(val)
	}
}

func TestMemory2(t *testing.T) {
	Memory().Set("key2", "val2", time.Second)
	if val, err := Memory().Get("key2"); err != nil {
		t.Error()
	} else {
		fmt.Println(val)
	}
	Memory().Delete("key2")
	if val, err := Memory().Get("key2"); err != nil {
		t.Error()
	} else {
		fmt.Println(val)
	}
}

func TestHash(t *testing.T) {
	hexists(t)
	hset(t)
	hget(t)
	hmset(t)
	hmget(t)
	hlen(t)

	hkeys(t)
	hvals(t)
	hgetall(t)

	hexists(t)
	hdel(t)
	hlen(t)
	hexists(t)
	hgetall(t)
}

const (
	key = "hashKey"
	field = "field"
)

func hdel(t *testing.T) {
	if err := Hash().HashDelete(key, field); err != nil {
		t.Error(err)
	} else {
		fmt.Println("hdel success")
	}
}

func hlen(t *testing.T) {
	if len, err := Hash().HashLen(key); err != nil {
		t.Error(err)
	} else {
		fmt.Printf("hlen success : %d\n", len)
	}
}

func hset(t *testing.T) {
	if err := Hash().HashSet(key, field, "valzczc"); err != nil {
		t.Error(err)
	} else {
		fmt.Println("hset success")
	}
}

func hget(t *testing.T) {
	if val, err := Hash().HashGet(key, field); err != nil {
		t.Error(err)
	} else {
		fmt.Printf("hget success : %s\n", val)
	}
}

func hmset(t *testing.T) {
	if err := Hash().HashMSet(key, field, "valzczczczxc"); err != nil {
		t.Error(err)
	} else {
		fmt.Println("hmset success")
	}
}

func hmget(t *testing.T) {
	if val, err := Hash().HashMGet(key, field); err != nil {
		t.Error(err)
	} else {
		fmt.Printf("hmget success : %+v\n", val)
	}
}

func hkeys(t *testing.T) {
	if keys, err := Hash().HashKeys(key); err != nil {
		t.Error(err)
	} else {
		fmt.Printf("hkeys success : %+v\n", keys)
	}
}

func hvals(t *testing.T) {
	if vals, err := Hash().HashVals(key); err != nil {
		t.Error(err)
	} else {
		fmt.Printf("hvals success : %+v\n", vals)
	}
}

func hgetall(t *testing.T) {
	if m, err := Hash().HashGetAll(key); err != nil {
		t.Error(err)
	} else {
		fmt.Printf("hgetall success : %+v\n", m)
	}
}

func hexists(t *testing.T) {
	if ok, err := Hash().HashExists(key, field); err != nil {
		t.Error(err)
	} else {
		fmt.Println(ok)
	}
}

