package lib

import (
  "database/sql"
  "github.com/joho/godotenv"
	"os"
  "fmt"
)

func LoadEnv() {
  godotenv.Load()
  godotenv.Load(fmt.Sprintf("/Users/d.kanai/workspace_go/src/spary_go/.env.%s", os.Getenv("GO_ENV")))
}

func DbOpen() *sql.DB {
  db, _ := sql.Open("mysql", os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@/" + os.Getenv("DB_NAME"))
  return db
}
