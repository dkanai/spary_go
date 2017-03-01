package batch

import (
  "fmt"
  "io/ioutil"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"lib"
)

func ImportOnsenList() {
  url := "http://jws.jalan.net/APICommon/OnsenSearch/V1/?key=aqr15a41839ced&l_area=010300&count=1&xml_ptn=1"
  req, _ := http.NewRequest("GET", url, nil)
  client := new(http.Client)
  resp, _ := client.Do(req)
  defer resp.Body.Close()

  byteArray, _ := ioutil.ReadAll(resp.Body)
  fmt.Println(string(byteArray))

	db := lib.Db_open()
	defer db.Close()

  query := "INSERT INTO spa (name,address) VALUES(?, ?)"

	rows, err := db.Query(query, "AAA", "Where")
	if err != nil {
		panic(err.Error())
	}
  fmt.Println(rows)
}
