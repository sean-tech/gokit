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

func NewHashStorage() IHashStorage {
	return new(hashStorageImpl)
}

type hashStorageImpl struct {
	hashStorageMap sync.Map
	len foundation.Int64
}

func (this *hashStorageImpl) newHashTable(key string) *sync.Map {
	var hashTable = new(sync.Map)
	this.hashStorageMap.Store(key, hashTable)
	return hashTable
}

func (this *hashStorageImpl) getHashTable(key string) (*sync.Map, bool) {
	hashInter, ok := this.hashStorageMap.Load(key)
	if ok == false {
		return nil, false
	}
	hashTable, ok := hashInter.(*sync.Map)
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
	_, ok = hashTable.Load(field)
	return ok, nil
}

func (this *hashStorageImpl) HashLen(key string) (int64, error) {
	return this.len.Load(), nil
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
		if _, ok = hashTable.Load(k); ok == false {
			this.len.Add(1)
		}
		hashTable.Store(k, v)
	}
	return nil
}

func (this *hashStorageImpl) HashGet(key, field string) (string, error) {
	hashTable, ok := this.getHashTable(key)
	if ok == false {
		return "", errors.New("hashtable in hashstorage : nil")
	}
	if valInter, ok := hashTable.Load(field); ok {
		return valInter.(string), nil
	} else {
		return "", errors.New("hashstorage : nil")
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
		if valInter, ok := hashTable.Load(field); ok {
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
		if _, ok := hashTable.Load(field); ok {
			hashTable.Delete(field)
			this.len.Sub(1)
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
	hashTable.Range(func(key, value interface{}) bool {
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
	hashTable.Range(func(key, value interface{}) bool {
		values = append(values, key.(string))
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
	hashTable.Range(func(key, value interface{}) bool {
		m[key.(string)] = value.(string)
		return true
	})
	return m, nil
}

