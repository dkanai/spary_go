package main

import (
	"fmt"
	"net/http"
)

func showOnsenList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Onsen List")
}

func showOnsenList2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Onsen List2")
}

func main() {
	http.HandleFunc("/onsen/list", showOnsenList)
	http.HandleFunc("/onsen/list2", showOnsenList2)
	http.ListenAndServe(":8080", nil)
}
