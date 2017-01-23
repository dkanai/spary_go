package main

import (
	"fmt"
	"net/http"
)

func showOnsenList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set( "Access-Control-Allow-Origin", r.RemoteAddr )
  w.Header().Set( "Access-Control-Allow-Credentials", "true" )
  w.Header().Set( "Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization" )
  w.Header().Set( "Access-Control-Allow-Methods","GET, POST, PUT, DELETE, OPTIONS" )
	fmt.Fprintf(w, "Onsen List")
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
	http.HandleFunc("/onsen/list", showOnsenList)
	http.HandleFunc("/chiiia12",showChiiia12)
	http.HandleFunc("/kika", showKika)
	fmt.Printf("Server is runnig... localhost:8080")
	http.ListenAndServe(":8080", nil)
}
