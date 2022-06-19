package handlers

import (
	"shoppinglist/src/data"
	"sync"
)

type ItemHandlers struct {
	sync.Mutex
  store map[string]data.Item
}
