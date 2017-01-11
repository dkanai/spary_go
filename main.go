package main

import (
	"fmt"
	"net/http"
)

func showOnsenList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Onsen List")
}

func showKika(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pc")
}

func main() {
	http.HandleFunc("/onsen/list", showOnsenList)
	http.HandleFunc("/kika", showKika)
	http.ListenAndServe(":8080", nil)
}
