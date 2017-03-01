package main

import (
  "api"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"os"
	"encoding/json"
  "lib"
  "database/sql"
)

var db *sql.DB

func TestMain(m *testing.M){
  lib.LoadEnv()

  db = lib.DbOpen()
  defer db.Close()

	code := m.Run()

  truncateTable()

	defer os.Exit(code)
}

func truncateTable() {
  db.Query("DELETE FROM spa;")
}

func execShowSpaList() *api.Result {
  //ShowSpaList実行
	req, _ := http.NewRequest("GET", "v1/spas", nil)
	res := httptest.NewRecorder()
	api.ShowSpaList(res, req)

  //レスポンスを構造体に変換
	data, _ := ioutil.ReadAll(res.Body)
	obj := new(api.Result)
	json.Unmarshal(([]byte)(string(data)), obj)
  return obj
}

func TestShowSpaList(t *testing.T) {
  db.Query("INSERT INTO spa (name, address) VALUES(?, ?)", "木下温泉", "北海道")
  db.Query("INSERT INTO spa (name, address) VALUES(?, ?)", "木下温泉2", "北海道2")

  result := execShowSpaList()

	if result.Spa[0].Name != "木下温泉" {
		t.Fatalf("%v != %v", string(result.Spa[0].Name), "木下温泉")
	}
	if result.Spa[0].Address != "北海道" {
		t.Fatalf("%v != %v", string(result.Spa[0].Address), "北海道")
	}
	if result.Spa[1].Name != "木下温泉2" {
		t.Fatalf("%v != %v", string(result.Spa[1].Name), "木下温泉2")
	}
	if result.Spa[1].Address != "北海道2" {
		t.Fatalf("%v != %v", string(result.Spa[1].Address), "北海道2")
	}
}
