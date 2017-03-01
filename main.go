package main

import (
  "api"
  "batch"
  "fmt"
  "os"
  "github.com/joho/godotenv"
)

func main() {
  godotenv.Load()
  fmt.Println(os.Getenv("ENV"))
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
