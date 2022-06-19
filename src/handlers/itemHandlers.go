package handlers

import (
	"net/http"
	"shoppinglist/src/data"
	"sync"
)

type ItemHandlers struct {
	sync.Mutex
  store map[string]data.Item
}

func (h * ItemHandlers) Test(w http.ResponseWriter, r *http.Request){
  w.Write([]byte("Hello world"))
}
