package storage

import (
	"sync"
	"time"
)

type IMemoryStorage interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) (string, error)
	Delete(key string)
}

var (
	_memoryStorage IMemoryStorage
	_memoryStorageOnce sync.Once
)
func Memory() IMemoryStorage {
	_memoryStorageOnce.Do(func() {
		_memoryStorage = NewMemeoryStorage()
	})
	return _memoryStorage
}

/**
* 获取内存存储实例
 */
func NewMemeoryStorage() IMemoryStorage {
	return new(memeoryStorageImpl)
}

// 内存存储实现
type memeoryStorageImpl struct {
	memoryStorageMap sync.Map
}

func (this *memeoryStorageImpl) Set(key string, value interface{}, expiresTime time.Duration) error {
	this.memoryStorageMap.Store(key, value)
	// 定时删除
	go func(storage *memeoryStorageImpl, expiresTime time.Duration) {
		select {
		case <- time.After(expiresTime):
			storage.Delete(key)
		}
	}(this, expiresTime)
	return nil
}

func (this *memeoryStorageImpl) Get(key string) (value string, err error) {
	if valInter, ok := this.memoryStorageMap.Load(key); ok {
		return valInter.(string), nil
	} else {
		return "", nil
	}
}

func (this *memeoryStorageImpl) Delete(key string) {
	this.memoryStorageMap.Delete(key)
}