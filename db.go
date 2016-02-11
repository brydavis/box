package main

// import (
// 	"database/sql"
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"

// 	//  _ "github.com/denisenkom/go-mssqldb"
// 	_ "./odbc"
// )

// var db *sql.DB

// func Connect(source string) {
// 	b, _ := ioutil.ReadFile(source)
// 	var config map[string]interface{}
// 	if err := json.Unmarshal(b, &config); err != nil {
// 		panic(err)
// 	}

// 	connString := fmt.Sprintf(
// 		"Dsn=%s;Driver={%s};uid=%s;pwd=%s;database=%s;host=%s;srvr=%s;serv=%s;pro=%s;cloc=%s;dloc=%s;opt=%s;",
// 		config["dsn"],
// 		config["driver-name"],
// 		config["user"],
// 		config["password"],
// 		config["database"],
// 		config["host"],
// 		config["server"],
// 		config["services"],
// 		config["protocol"],
// 		config["cloc"],
// 		config["dloc"],
// 		config["opt"],
// 	)

// 	db, err := sql.Open(config["driver-type"].(string), connString)
// 	if err != nil {
// 		log.Panic("Open connection failed:", err.Error())
// 	}
// 	// defer db.Close()

// 	if err = db.Ping(); err != nil {
// 		log.Panic(err)
// 	} else {
// 		fmt.Println("Successfully connected to the database")
// 	}
// }
