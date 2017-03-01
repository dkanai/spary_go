package batch

import (
	"database/sql"
  "fmt"
  "io/ioutil"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"os"
)

func ImportOnsenList() {
  // バッチ処理
  url := "http://jws.jalan.net/APICommon/OnsenSearch/V1/?key=aqr15a41839ced&l_area=010300&count=1&xml_ptn=1"
  req, _ := http.NewRequest("GET", url, nil)
  client := new(http.Client)
  resp, _ := client.Do(req)
  defer resp.Body.Close()

  byteArray, _ := ioutil.ReadAll(resp.Body)
  fmt.Println(string(byteArray))

	db, err := sql.Open("mysql", os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@/" + os.Getenv("DB_NAME"))

	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // 関数がリターンする直前に呼び出される

  query := "INSERT INTO spa (name,address) VALUES(?, ?)"

	rows, err := db.Query(query, "AAA", "Where") //
	if err != nil {
		panic(err.Error())
	}
  fmt.Println(rows)
}
