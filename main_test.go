package main

import (
  "api"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"os"
	"encoding/json"
  "github.com/joho/godotenv"
  "lib"
)

func TestMain(m *testing.M){
  godotenv.Load()
  godotenv.Load(fmt.Sprintf(".env.%s", os.Getenv("GO_ENV")))
	code := m.Run()
	defer os.Exit(code)
}

func TestShowSpaList(t *testing.T) {
  db := lib.Db_open()
  defer db.Close()
  query := "DELETE FROM spa;"
  db.Query(query)
  query2 := "INSERT INTO spa (name,address) VALUES(?, ?)"
  db.Query(query2, "木下温泉", "木下温泉")

	var requests [3]*http.Request
	var err error

	requests[0], err = http.NewRequest("GET", "/api/spa/list", nil)
	if err != nil {
		t.Errorf("NewRequest[0] Error. %v", err)
	}

	r := httptest.NewRecorder()

	api.ShowSpaList(r, requests[0])

	data, err := ioutil.ReadAll(r.Body)
	obj := new(api.Result)
	json.Unmarshal(([]byte)(string(data)),obj)
	if obj.Spa[0].Name != "木下温泉" {
		t.Fatalf("Data Error. name is not '木下温泉'. %v",string(obj.Spa[0].Name))
	 }
}
