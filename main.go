package main

import (
	"database/sql"
	"fmt"

	"github.com/MakeNowJust/heredoc/v2"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./sample.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("DB is opened.")

	cursor, err := db.Query("SELECT * FROM Example;")
	if err != nil {
		panic(err)
	}
	defer cursor.Close()
	fmt.Println("Cursor is opened.")

	debugFmt := heredoc.Doc(`
		- - - - - - - - - - -
		columnA → %s
		columnB → %s
		columnC → %s
		columnD → %s
		= = = = = = = = = = =
	`)
	for cursor.Next() {
		var columnA string
		var columnB string
		var columnC string
		var columnD string
		err := cursor.Scan(
			&columnA,
			&columnB,
			&columnC,
			&columnD,
		)
		fmt.Printf(
			debugFmt,
			columnA,
			columnB,
			columnC,
			columnD,
		)
		if err != nil {
			panic(err)
		}

	}
	err = cursor.Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("Finito")

}
