package main

import (
	"os"
	"shoppinglist/src"
)

var port string = os.Getenv("SERVER_PORT")

func main() {

  if port == "" {
    port = "8081"
  }
  
  src.StartServer(port)
}
