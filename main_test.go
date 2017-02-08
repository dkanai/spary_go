package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
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
	//	requests[0].Header.Add("User-Agent", "Mozilla/5.0 (iPad; CPU OS 8_1_2 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) CriOS/39.0.2171.50 Mobile/12B440 Safari/600.1.4")

	requests[1], err = http.NewRequest("GET", "/api/spa/list", nil)
	if err != nil {
		t.Errorf("NewRequest[1] Error. %v", err)
	}
	//	requests[1].Header.Add("User-Agent", "Mozilla/5.0 (Linux; Android 4.4.2; 302KC Build/101.0.2c00) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/30.0.0.0 Mobile Safari/537.36")

	for pos, req := range requests {
		r := httptest.NewRecorder()

		showSpaList(r, req)

		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("Error by ioutil.ReadAll(). %v", err)
		}

		if r.Code == 200 {
			if "Hello HTTP Test" != string(data) {
				t.Fatalf("Data Error. %v", string(data))
			}
		} else {
			if r.Code != 404 {
				t.Fatalf("Status Error %d", r.Code)
			}
			expect := "{\"spa\":[{\"name\":\"木下温泉\",\"address\":\"東日本橋\"},{\"name\":\"hoge温泉\",\"address\":\"大坂\"}]}"

			if string(expect) != string(data) {
				t.Fatalf("Data Error. %v", string(data))
			}

			if pos != 0 {
				t.Fatalf("Request Error. %d", pos)
			}
		}
	}
}
