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
  "database/sql"
)

var sampleHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello HTTP Test")
})

func TestSampleHandler(t *testing.T) {
	ts := httptest.NewServer(sampleHandler)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/api/spa/list")
	if err != nil {
		t.Error("unexpected")
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Error by ioutil.ReadAll(). %v", err)
	}

	if "Hello HTTP Test" != string(data) {
		t.Fatalf("Data Error. %v", string(data))
	}
}
func TestMain(m *testing.M){
  godotenv.Load()
  godotenv.Load(fmt.Sprintf(".env.%s", os.Getenv("GO_ENV")))
	code := m.Run()
	defer os.Exit(code)
}

func TestShowSpaList(t *testing.T) {
  db, _ := sql.Open("mysql", os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@/" + os.Getenv("DB_NAME"))
  defer db.Close()
  query := "DELETE FROM spa;"
  db.Query(query, "木下温泉", "Where")
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
