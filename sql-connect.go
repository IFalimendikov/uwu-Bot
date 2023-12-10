package main

import (
	"database/sql"
	"log"
	"os"
	"time"
	"strings"
	// "math/rand"

	_ "github.com/go-sql-driver/mysql"
)

func SqlConnect(id uint64) (string, string, int) {
	var randomUwu string
	var artistLink string
	var returnedId int

	sqlToken := os.Getenv("SQL_TOKEN")

	db, err := sql.Open("mysql", sqlToken)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	if id == 0 {
		rows1, err := db.Query("SELECT imageLink, socialMedia, uwucrewId FROM uwuDerivatives ORDER BY RAND() LIMIT 1")
		if err != nil {
			panic(err)
		}

		for rows1.Next() {
			var column1Value string
			var column2Value string
			var column3Value int
			err := rows1.Scan(&column1Value, &column2Value, &column3Value)
			if err != nil {
				panic(err)
			}
			randomUwu = column1Value
			artistLink = column2Value
			returnedId = column3Value

			if !strings.HasPrefix(artistLink, "https://") {
				artistLink = "https://" + artistLink
			}
		}
		
		err = rows1.Close()
		if err != nil {
			log.Fatal(err)
		}

	} else {
		rows1, err := db.Query("SELECT imageLink, socialMedia, uwucrewId FROM uwuDerivatives WHERE uwucrewId =  ? ORDER BY RAND()", id)
		if err != nil {
			panic(err)
		}

		if rows1.Next() {
			for rows1.Next() {
				var column1Value string
				var column2Value string
				var column3Value int
				err := rows1.Scan(&column1Value, &column2Value, &column3Value)
				if err != nil {
					panic(err)
				}
				randomUwu = column1Value
				artistLink = column2Value
				returnedId = column3Value

				if !strings.HasPrefix(artistLink, "https://") {
					artistLink = "https://" + artistLink
				}
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
	}

	return randomUwu, artistLink, returnedId
}
