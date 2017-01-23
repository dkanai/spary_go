package main

import (
	"fmt"
	"net/http"
)

func showOnsenList(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Onsen List From Go API")
}

func showOnsenList2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Onsen List")
}

func showKika(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pc")
}
func showChiiia12(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "chiiia12")
}

func main() {
	http.HandleFunc("/api/onsen/list", showOnsenList)
	http.HandleFunc("/chiiia12",showChiiia12)
	http.HandleFunc("/kika", showKika)
	fmt.Printf("Server is runnig... localhost:8080")
	http.ListenAndServe(":8080", nil)
}
