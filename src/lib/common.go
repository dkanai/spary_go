package lib

import (
  "database/sql"
	"os"
)

func Db_open() *sql.DB {
  db, _ := sql.Open("mysql", os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@/" + os.Getenv("DB_NAME"))
  return db
}
