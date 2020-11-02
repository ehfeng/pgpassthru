package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type QueryRequest struct {
	Dsn      string        `json:"dsn"`
	Sql      string        `json:"sql"`
	Bindings []interface{} `json:"bindings"`
}

type QueryResponseCol struct {
	Name     string `json:"name"`
	Datatype string `json:"datatype"`
}

type StatementType string

const (
	StatementTypeSelect  = "select"
	StatementTypeInsert  = "insert"
	StatementTypeDelete  = "delete"
	StatementTypeUpdate  = "update"
	StatementTypeUnknown = "unknown"
)

type QueryResponse struct {
	Result struct {
		Cols []QueryResponseCol `json:"cols"`
		Rows [][]string         `json:"rows"`
	} `json:"result"`
	Error string `json:"error"`
}

func query(w http.ResponseWriter, r *http.Request) {
	errorResponse := func(err error) []byte {
		b, err := json.Marshal(QueryResponse{Error: err.Error()})
		if err != nil {
			panic(err)
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(b)
		return b
	}

	if r.Method != "POST" {
		log.Println("Only POST method accepted")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Bad json")
		errorResponse(err)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		log.Println("Requires Content-Type application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	queryRequest := QueryRequest{}
	if err = json.Unmarshal(bodyBytes, &queryRequest); err != nil {
		log.Println("Unable to unmarshal request")
		errorResponse(err)
		return
	}

	conn, err := sql.Open("pgx", queryRequest.Dsn)
	if err != nil {
		log.Println("Unable to connect with dsn", queryRequest.Dsn)
		errorResponse(err)
		return
	}
	defer conn.Close()
	rows, err := conn.Query(queryRequest.Sql, queryRequest.Bindings...)
	if err != nil {
		log.Println("Query error")
		errorResponse(err)
		return
	}
	defer rows.Close()

	resp := QueryResponse{}
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		log.Println("ColumnTypes failed")
		errorResponse(err)
		return
	}
	columnNames, err := rows.Columns()
	columnCount := len(columnNames)
	if err != nil {
		errorResponse(err)
		return
	}
	for i, columnName := range columnNames {
		resp.Result.Cols = append(resp.Result.Cols, QueryResponseCol{Name: columnName, Datatype: columnTypes[i].ScanType().Name()})
	}
	scanRow := make([]*string, columnCount)
	scanRowPointers := make([]interface{}, columnCount)
	for i := 0; i < columnCount; i++ {
		scanRowPointers[i] = &scanRow[i]
	}

	for rows.Next() {
		if err = rows.Scan(scanRowPointers...); err != nil {
			log.Println("Row scan failed", err)
			errorResponse(err)
			return
		}
		record := make([]string, columnCount)
		for i, value := range scanRow {
			if value == nil {
				record[i] = ""
			} else {
				record[i] = *scanRow[i]
			}
		}
		resp.Result.Rows = append(resp.Result.Rows, record)
	}

	respBytes, err := json.Marshal(resp)
	if err != nil {
		errorResponse(err)
		return
	}
	if _, err = w.Write(respBytes); err != nil {
		errorResponse(err)
		return
	}

}

func main() {
	http.HandleFunc("/", query)
	fmt.Println("Listening on http://locahost:8090")
	http.ListenAndServe(":8090", nil)
}
