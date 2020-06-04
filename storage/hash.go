package storage

import (
	"errors"
	"github.com/sean-tech/gokit/foundation"
	"sync"
)

type IHashStorage interface {
	// hash table set & get
	HashExists(key, field string) (bool, error)
	HashLen(key string) (int64, error)
	HashSet(key string, values ...interface{}) error
	HashGet(key, field string) (string, error)
	HashMSet(key string, values ...interface{}) error
	HashMGet(key string, fields ...string) ([]interface{}, error)
	HashDelete(key string, fields ...string) error
	HashKeys(key string) ([]string, error)
	HashVals(key string) ([]string, error)
	HashGetAll(key string) (map[string]string, error)
}

var (
	_hashStorage IHashStorage
	_hashStorageOnce sync.Once
)

func Hash() IHashStorage {
	_hashStorageOnce.Do(func() {
		_hashStorage = NewHashStorage()
	})
	return _hashStorage
}

type hashTable struct {
	hashMap sync.Map
	len *foundation.Int64
}
func NewHashStorage() IHashStorage {
	return new(hashStorageImpl)
}

type hashStorageImpl struct {
	hashStorageMap sync.Map
}

func (this *hashStorageImpl) newHashTable(key string) *hashTable {
	var hashTable = &hashTable{
		hashMap: sync.Map{},
		len:     foundation.NewInt64(0),
	}
	this.hashStorageMap.Store(key, hashTable)
	return hashTable
}

func (this *hashStorageImpl) getHashTable(key string) (*hashTable, bool) {
	hashInter, ok := this.hashStorageMap.Load(key)
	if ok == false {
		return nil, false
	}
	hashTable, ok := hashInter.(*hashTable)
	if ok == false {
		return nil, false
	}
	return hashTable, true
}


func (this *hashStorageImpl) HashExists(key, field string) (bool, error) {
	hashTable, ok := this.getHashTable(key)
	if ok == false {
		return ok, nil
	}
	_, ok = hashTable.hashMap.Load(field)
	return ok, nil
}

func (this *hashStorageImpl) HashLen(key string) (int64, error) {
	hashTable, ok := this.getHashTable(key)
	if ok == false {
		return 0, nil
	}
	return hashTable.len.Load(), nil
}

func (this *hashStorageImpl) HashSet(key string, values ...interface{}) error {
	hashTable, ok := this.getHashTable(key)
	if ok == false {
		hashTable = this.newHashTable(key)
	}
	if len(values) % 2 != 0  {
		return errors.New("wrong number of arguments for hashset in hashstorage")
	}
	for idx := 1; idx < len(values); idx ++ {
		k := values[idx-1]; v := values[idx]
		if _, ok = hashTable.hashMap.Load(k); ok == false {
			hashTable.len.Add(1)
		}
		hashTable.hashMap.Store(k, v)
	}
	return nil
}

func (this *hashStorageImpl) HashGet(key, field string) (string, error) {
	hashTable, ok := this.getHashTable(key)
	if ok == false {
		return "", errors.New("hashtable in hashstorage : nil")
	}
	value, ok := hashTable.hashMap.Load(field)
	if ok == false {
		return "", errors.New("hashstorage : nil")
	}
	switch value.(type) {
	case string:
		return value.(string), nil
	case byte:
		return string(value.(byte)), nil
	case []byte:
		return string(value.([]byte)), nil
	default:
		return "", nil
	}
}

func (this *hashStorageImpl) HashMSet(key string, values ...interface{}) error {
	return this.HashSet(key, values...)
}

func (this *hashStorageImpl) HashMGet(key string, fields ...string) ([]interface{}, error) {
	var values []interface{}

	hashTable, ok := this.getHashTable(key)
	if ok == false {
		return values, nil
	}
	for _, field := range fields {
		if valInter, ok := hashTable.hashMap.Load(field); ok {
			values = append(values, valInter)
		} else {
			values = append(values, "")
		}
	}
	return values, nil
}

func (this *hashStorageImpl) HashDelete(key string, fields ...string) error {
	hashTable, ok := this.getHashTable(key)
	if ok == false {
		return nil
	}
	for _, field := range fields {
		if _, ok := hashTable.hashMap.Load(field); ok {
			hashTable.hashMap.Delete(field)
			hashTable.len.Sub(1)
		}
	}
	return nil
}

func (this *hashStorageImpl) HashKeys(key string) ([]string, error) {
	var keys []string
	hashTable, ok := this.getHashTable(key)
	if ok == false {
		return keys, nil
	}
	hashTable.hashMap.Range(func(key, value interface{}) bool {
		keys = append(keys, key.(string))
		return true
	})
	return keys, nil
}

func (this *hashStorageImpl) HashVals(key string) ([]string, error) {
	var values []string
	hashTable, ok := this.getHashTable(key)
	if ok == false {
		return values, nil
	}
	hashTable.hashMap.Range(func(key, value interface{}) bool {
		values = append(values, value.(string))
		return true
	})
	return values, nil
}

func (this *hashStorageImpl) HashGetAll(key string) (map[string]string, error) {
	var m = make(map[string]string)
	hashTable, ok := this.getHashTable(key)
	if ok == false {
		return m, nil
	}
	hashTable.hashMap.Range(func(key, value interface{}) bool {
		m[key.(string)] = value.(string)
		return true
	})
	return m, nil
}

