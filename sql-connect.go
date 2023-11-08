package main

import (

	"log"
	"database/sql"
	"time"
	"os"
	// "math/rand"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func SqlConnect (id uint64) (string, string) {
	var randomUwu string
	var artistLink string

	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
	sqlToken := os.Getenv("SQL_TOKEN")

	db, err := sql.Open("mysql", sqlToken)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)



	rows1, err := db.Query("SELECT imageLink, socialMedia FROM uwuDerivatives WHERE uwucrewId =  ? ORDER BY RAND()", id)
	if err != nil {
		panic(err)
	}

	if rows1.Next() {
		for rows1.Next() {
			var column1Value string
			var column2Value string
			err := rows1.Scan(&column1Value, &column2Value)
			if err != nil {
				panic(err)
			}
			randomUwu = column1Value
			artistLink = column2Value
		}

		// rand.Seed(time.Now().UnixNano())	

		// randomIndex := rand.Intn(len(results))
		// randomUwuPic := results[randomIndex]

		// rows2, err := db.Query("SELECT socialMedia FROM uwuDerivatives WHERE imageLink = ?", randomUwuPic)
		// if err != nil {
		// 	panic(err)
		// }

		// artist = rows 

		// for rows2.Next() {
		// 	var column2Value string
		// 	err := rows1.Scan(&column2Value)
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	results = append(results, column2Value)
		// }

	} else {
		randomUwu = "no ID found"
	}

	err = rows1.Close()
	if err != nil {
		log.Fatal(err)
	}

	return randomUwu, artistLink
}