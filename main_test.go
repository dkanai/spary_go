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

  truncateTable()

	code := m.Run()

	defer os.Exit(code)
}

func truncateTable() {
  query := "TRUNCATE FROM spa;"
  db.Query(query)
}

func TestShowSpaList(t *testing.T) {
  query := "INSERT INTO spa (name,address) VALUES(?, ?)"
  db.Query(query, "木下温泉", "木下温泉")

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
