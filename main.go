package main

import (
	"os"
	"shoppinglist/src"
)

var port string = os.Getenv("SERVER_PORT") || 8081

func main() {
  src.StartServer()
}
