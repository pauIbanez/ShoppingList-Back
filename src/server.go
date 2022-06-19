package src

import (
	"net/http"
	"shoppinglist/src/handlers"
)

func StartServer(port string) {
  itemhandlers := handlers.CreateItemHandler()

  http.HandleFunc("/items/all", itemhandlers.GetAllItems)

  err := http.ListenAndServe(":"+port, nil)
  if err != nil {
    panic(err)
  }
}
