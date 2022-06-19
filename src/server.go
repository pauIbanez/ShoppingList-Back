package src

import (
	"net/http"
	"shoppinglist/src/handlers"
)

func StartServer(port string) {
  itemhandlers := handlers.CreateItemHandler()

  http.HandleFunc("/items/all", itemhandlers.GetAllItems)
  http.HandleFunc("/items/modify", itemhandlers.ModifyItem)
  http.HandleFunc("/items/create", itemhandlers.CreateItem)
  http.HandleFunc("/items/delete", itemhandlers.DeleteItem)

  err := http.ListenAndServe(":"+port, nil)
  if err != nil {
    panic(err)
  }
}
