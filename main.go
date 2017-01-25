package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func showOnsenList(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:@/sample")
	//	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // 関数がリターンする直前に呼び出される

	rows, err := db.Query("SELECT * FROM sample_table") //
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

	for rows.Next() {
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
			fmt.Fprintf(w,columns[i], ": ", value)
		}
		fmt.Fprintf(w,"-----------------------------------\n")
	}
}

func main() {
	http.HandleFunc("/spa/list", showOnsenList)
//	fmt.Printf("Server is runnig... localhost:8080")
	http.ListenAndServe(":8080", nil)

}
