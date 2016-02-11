package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"

	// _ "github.com/denisenkom/go-mssqldb"
	// _ "./odbc"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Connect(source string) {
	b, err := ioutil.ReadFile(source)
	if err != nil {
		fmt.Println(err)
	}
	var config map[string]interface{}
	if err := json.Unmarshal(b, &config); err != nil {
		panic(err)
	}

	if driverType := config["driver-type"].(string); driverType == "sqlite3" {
		db, err = sql.Open(driverType, config["database"].(string))
		if err != nil {
			log.Panic("Open connection failed:", err.Error())
		}
	} else {
		var connString string
		for k, v := range config {
			connString += fmt.Sprintf("%s=%v;", k, v)
		}

		db, err = sql.Open(config["driver-type"].(string), connString)
		if err != nil {
			log.Panic("Open connection failed:", err.Error())
		}
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}
}

func ListenAndServe(port int) error {
	http.HandleFunc("/static/", StaticHandler)
	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/metrics", MetricsHandler)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		return err
	}
	return nil
}

func StaticHandler(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, req.URL.Path[1:])
}

func RootHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		if err := db.Ping(); err != nil {
			log.Panic(err)
		} else {
			fmt.Println("Successfully connected to the database")
		}

		b, _ := ioutil.ReadFile("views/index.html")
		t := template.New("")
		t, _ = t.Parse(string(b))
		t.Execute(res, nil)
	case "POST":
		req.ParseForm()
		query := req.FormValue("code") // add support for multiple queries
		rows, err := db.Query(query)
		if err != nil {
			fmt.Println("HERE\n")
			fmt.Println(err.Error(), "\n")

		}
		defer rows.Close()

		columns, _ := rows.Columns()
		count := len(columns)
		values := make([]interface{}, count)
		valuePtrs := make([]interface{}, count)

		var dataset []map[string]interface{}

		for rows.Next() {
			for i, _ := range columns {
				valuePtrs[i] = &values[i]
			}

			rows.Scan(valuePtrs...)
			datarow := make(map[string]interface{})
			for i, col := range columns {
				var v interface{}
				val := values[i]
				b, ok := val.([]byte)

				if ok {
					v = string(b)
				} else {
					v = val
				}
				datarow[col] = v
			}
			dataset = append(dataset, datarow)
		}
		results, _ := json.Marshal(dataset)
		payload, _ := json.Marshal(map[string]string{"code": query, "output": string(results)})
		res.Write(payload)
	}
}

func MetricsHandler(res http.ResponseWriter, req *http.Request) {

	b, _ := ioutil.ReadFile("views/metrics_test.html")
	t := template.New("")
	t, _ = t.Parse(string(b))
	t.Execute(res, nil)
}
