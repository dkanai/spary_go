package main

import (
  "fmt"
  "io/ioutil"
	"net/http"
)

func main() {
  // バッチ処理
  url := "http://jws.jalan.net/APICommon/OnsenSearch/V1/?key=aqr15a41839ced&l_area=010300&count=1&xml_ptn=1"
  req, _ := http.NewRequest("GET", url, nil)
  client := new(http.Client)
  resp, _ := client.Do(req)
  defer resp.Body.Close()

  byteArray, _ := ioutil.ReadAll(resp.Body)
  fmt.Println(string(byteArray))
}
