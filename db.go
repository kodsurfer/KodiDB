package main

import "sync"

type DB struct {
  lock sync.RWMutex
}
