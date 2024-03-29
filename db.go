package main

import (
	os "os"
	"sync"
)

type keyLocks struct {
	sync.RWMutex
	keys map[uint64]struct{}
}

func (k *keyLocks) add(key uint64) {
	k.Lock()
	defer k.Unlock()
	k.keys[key] = struct{}{}
}

func (k *keyLocks) exist(key uint64) bool {
	k.RLock()
	defer k.RUnlock()
	_, ok := k.keys[key]
	return ok
}

func (k *keyLocks) list() []uint64 {
	k.RLock()
	defer k.RUnlock()
	keys := make([]uint64, 0, len(k.keys))
	for key := range k.keys {
		keys = append(keys, key)
	}
	return keys
}

type DB struct {
	lock sync.RWMutex
	data *data
}

func Open(filepath string) (*DB, error) {
	pagesize := os.Getpagesize()

	newData, err := NewData(filepath, pagesize)
	if err != nil {
		return nil, err
	}

	db := &DB{
		lock: sync.RWMutex{},
		data: newData,
	}

	return db, nil
}
