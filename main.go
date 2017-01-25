package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)
type Spa struct {
	Name    string
	Address string
}

func showSpaList(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:@/spary")
	//	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // 関数がリターンする直前に呼び出される

	rows, err := db.Query("SELECT * FROM spa") //
	if err != nil {
		panic(err.Error())
	}

	columns, err := rows.Columns() // カラム名を取得
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))
	//  rows.Scan は引数に `[]interface{}`が必要.

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	spaList :=[]Spa{}
	for rows.Next() {
		spa := Spa{}
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			if columns[i] == "name" {
				spa.Name = value
			} else {
				spa.Address = value
			}
		}
		spaList = append(spaList,spa)

	}
	data :=spaList 
	bytes, err := json.Marshal(data)
	if err != nil {
		return
	}
	fmt.Fprintf(w, "%s\n", string(bytes))
}

func main() {
	http.HandleFunc("/api/spa/list", showSpaList)
	//	fmt.Printf("Server is runnig... localhost:8080")
	http.ListenAndServe(":8080", nil)

}

