package main

import (

	"log"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func SqlConnect (id uint64) []string {
	var results []string

	db, err := sql.Open("mysql", "cat:uwu@tcp(34.32.9.223:3306)/uwu")
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	rows, err := db.Query("SELECT imageLink FROM uwuDerivatives WHERE uwucrewId =  ?", id)
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		for rows.Next() {
			var column2Value string
			err := rows.Scan(&column2Value)
			if err != nil {
				panic(err)
			}
			results = append(results, column2Value)

			err = rows.Close()
			if err != nil {
				log.Fatal(err)
			}
		}
	} else {
		results = append(results, "no ID found")
	}
	return results
}