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
  godotenv.Load(fmt.Sprintf(".env.%s", os.Getenv("GO_ENV")))
  fmt.Println(os.Getenv("DB_NAME"))
  fmt.Println(os.Args)

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
