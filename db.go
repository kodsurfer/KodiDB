package main

import "sync"

type keyLocks struct {
  sync.RWMutex
  keys map[uint64]struct{}
}

func (k *keyLocks) add(key uint64) {
  k.Lock()
  defer k.Unlock()
  k.keys[key] = struct{}{}
}

type DB struct {
  lock sync.RWMutex
}
