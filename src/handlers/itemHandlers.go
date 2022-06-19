package handlers

import (
	"encoding/json"
	"net/http"
	"shoppinglist/src/data"
	"sync"
)

type ItemHandlers struct {
	sync.Mutex
  store map[string]data.Item
}

func (h * ItemHandlers) GetAllItems(w http.ResponseWriter, r *http.Request){

  items := make([]data.Item, len(h.store))

  i := 0
  h.Lock()
    for _, item := range h.store {
    items[i] = item
    i++
  }
  h.Unlock()

  jsonBytes, err := json.Marshal(items)

  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(err.Error()))
    return
  }

  w.Header().Add("content-type", "application/json")
  w.Write(jsonBytes)
}
