package api

import (
	"fmt"
	"net/http"
)

func Run() {
	http.HandleFunc("/v1/spas", ShowSpaList)

	fmt.Printf("Server is running... localhost:8080")
	http.ListenAndServe(":8080", nil)
}
