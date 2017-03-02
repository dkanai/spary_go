package api

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"lib"
)

type Result struct {
	Spas []Spa `json:"spa"`
}

type Spa struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

func ShowSpaList(w http.ResponseWriter, r *http.Request) {
	db := lib.DbOpen()
	defer db.Close()

	rows, _ := db.Query("SELECT * FROM spa")
	result := Result{}
	for rows.Next() {
		spa := Spa{}
		rows.Scan(&spa.Id, &spa.Name, &spa.Address)
		result.Spas = append(result.Spas, spa)
	}
	bytes, _ := json.Marshal(result)

	fmt.Fprintf(w, "%s", string(bytes))
}
