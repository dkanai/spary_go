package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
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
func TestShowSpaList(t *testing.T) {
	var requests [2]*http.Request
	var err error

	requests[0], err = http.NewRequest("GET", "/api/spa/list", nil)
	if err != nil {
		t.Errorf("NewRequest[0] Error. %v", err)
	}

	r := httptest.NewRecorder()

	showSpaList(r, requests[0])

	data, err := ioutil.ReadAll(r.Body)
	obj := new(Result)
	json.Unmarshal(([]byte)(string(data)),obj)
	if obj.Spa[0].Name != "木下温泉" {
		t.Fatalf("Data Error. name is not '木下温泉'. %v",string(obj.Spa[0].Name))
	 }
	if obj.Spa[0].Address != "東日本橋" {
		t.Fatalf("Data Error. address is not '東日本橋'. %v",string(obj.Spa[0].Address))
	}
	if obj.Spa[1].Name!= "hoge温泉" {
		t.Fatalf("Data Error. name is not 'hoge温泉'. %v",string(obj.Spa[1].Name))
	}
	if obj.Spa[1].Address != "大坂" {
		t.Fatalf("Data Error. address is not '大坂'. %v",string(obj.Spa[1].Address))
	}
}
