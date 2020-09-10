package storage

import (
	"errors"
	"sync"
)

var kvDB *KvDB
var once sync.Once

type LyMap struct {
	sync.RWMutex
	m map[string]*KvObject
}

type KvDB struct {
	dict LyMap
}

func NewKvDB() *KvDB {
	kvDB := &KvDB{
		dict: LyMap{
			m: make(map[string]*KvObject),
		},
	}
	return kvDB
}

func GetKvDb() *KvDB {
	once.Do(func() {
		kvDB = NewKvDB()
	})
	return kvDB
}

func (k *KvDB) dbAdd(key string, object *KvObject) error {
	if object == nil {
		return errors.New("object is nil")
	}
	defer k.dict.Unlock()
	k.dict.Lock()
	k.dict.m[key] = object
	return nil
}

func (k *KvDB) dbFind(key string) (*KvObject, error) {
	defer k.dict.RUnlock()
	k.dict.RLock()
	if object, ok := k.dict.m[key]; ok {
		return object, nil
	}
	return nil, errors.New("nil")
}

func (k *KvDB) dbDelete(key string) {
	defer k.dict.Unlock()
	k.dict.Lock()
	if _, ok := k.dict.m[key]; ok {
		delete(k.dict.m, key)
	}
}

func (k *KvDB) SetKey(key string, object *KvObject) error {
	if err := k.dbAdd(key, object); err != nil {
		return err
	}
	return nil
}

func (k *KvDB) DeleteKey(key string) {
	k.dbDelete(key)
}

func (k *KvDB) FindKey(key string) (*KvObject, error) {
	object, err := k.dbFind(key)
	return object, err
}
