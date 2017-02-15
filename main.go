package main

import (
  "api"
  "batch"
  "fmt"
  "os"
)

func main() {
  name := os.Args[1]
  fmt.Println(name)
  if name == "api" {
    fmt.Println("run router")
    api.Run()
  } else if name == "batch" {
    fmt.Println("run router")
    batch.ImportOnsenList()
  }
}
