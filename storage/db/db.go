package db

import (
	"errors"
	"github.com/ldaysjun/lykv/storage"
)

type KvDB struct {
	dict map[string]*storage.KvObject
}

func NewKvDB() *KvDB {
	kvDB := &KvDB{
		dict:make(map[string]*storage.KvObject),
	}
	return kvDB
}

func (k *KvDB) dbAdd(key string, object *storage.KvObject) error {
	if object == nil {
		return errors.New("object is nil")
	}
	k.dict[key] = object
	return nil
}

func (k *KvDB) dbFind(key string) *storage.KvObject {
	if object, ok := k.dict[key]; ok {
		return object
	}
	return nil
}

func (k *KvDB) dbDelete(key string) {
	if _, ok := k.dict[key]; ok {
		delete(k.dict, key)
	}
}

func (k *KvDB) SetKey(key string, object *storage.KvObject) error {
	if err := k.dbAdd(key, object); err != nil {
		return err
	}
	return nil
}

func (k *KvDB) DeleteKey(key string) {
	k.dbDelete(key)
}

func (k *KvDB) FindKey(key string) *storage.KvObject {
	object := k.dbFind(key)
	return object
}
