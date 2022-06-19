package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"shoppinglist/src/data"
	"strings"
	"sync"
	"time"
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

func (h * ItemHandlers) ModifyItem(w http.ResponseWriter, r *http.Request){


  bodyBytes, err := ioutil.ReadAll(r.Body)
  defer r.Body.Close()

  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(err.Error()))
    return
  }

  ct := r.Header.Get("content-type")

  if !strings.Contains(ct, "application/json")  {
    w.WriteHeader(http.StatusUnsupportedMediaType)
    w.Write([]byte("Only json is supported"))
    return
  }

  var item data.Item

  err = json.Unmarshal(bodyBytes, &item)

  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte(err.Error()))
    return
  }

  if item.Id == "" {
     w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte("Bad request, item Id needed"))
    return
  }

  h.Lock()
  foundItem, ok := h.store[item.Id]
  h.Unlock()

  if !ok {
    w.WriteHeader(http.StatusNotFound)
    return
  }

  if item.Name == "" {
    item.Name = foundItem.Name
  }
  if item.Quantity == 0 {
    item.Quantity = foundItem.Quantity
  }


  h.Lock()
  h.store[item.Id] = item
  h.Unlock()

  w.WriteHeader(http.StatusAccepted)
}

func (h * ItemHandlers) CreateItem(w http.ResponseWriter, r *http.Request) {
  bodyBytes, err := ioutil.ReadAll(r.Body)
  defer r.Body.Close()
  
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(err.Error()))
    return
  }

  ct := r.Header.Get("content-type")

  if !strings.Contains(ct, "application/json")  {
    w.WriteHeader(http.StatusUnsupportedMediaType)
    w.Write([]byte("Only json is supported"))
    return
  }

  var item data.Item
  err = json.Unmarshal(bodyBytes, &item)

  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte(err.Error()))
    return
  }

  if item.Quantity == 0 || item.Name == "" {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte("Invalid Request"))
    return
  }

  item.Id = fmt.Sprintf("%d", time.Now().UnixNano())

  h.Lock()
  h.store[item.Id] = item
  defer h.Unlock()
  
  w.WriteHeader(http.StatusCreated)
}

func (h * ItemHandlers) DeleteItem(w http.ResponseWriter, r *http.Request) {
  bodyBytes, err := ioutil.ReadAll(r.Body)
  defer r.Body.Close()
  
  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte(err.Error()))
    return
  }

  ct := r.Header.Get("content-type")

  if !strings.Contains(ct, "application/json")  {
    w.WriteHeader(http.StatusUnsupportedMediaType)
    w.Write([]byte("Only json is supported"))
    return
  }

  var item data.Item
  err = json.Unmarshal(bodyBytes, &item)

  if err != nil {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte(err.Error()))
    return
  }

  if item.Id == "" {
     w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte("Bad request, item Id needed"))
    return
  }

  var Id = item.Id 


  h.Lock()
  delete(h.store, Id)
  defer h.Unlock()
}





