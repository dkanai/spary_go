package main

import (
  "api"
  "batch"
  "fmt"
  "os"
  "lib"
)

func main() {
  lib.LoadEnv()

  if len(os.Args) == 1 {
    api.Run()
  } else if os.Args[1] == "api" {
    api.Run()
  } else if os.Args[1] == "batch" {
    batch.ImportOnsenList()
  } else {
    fmt.Println("not work")
  }
}
