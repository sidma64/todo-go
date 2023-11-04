package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
)

func main() {
	connStr := "postgresql://sidma64:0eNEqKJwtY6p@ep-withered-voice-43063416.eu-central-1.aws.neon.tech/neondb"
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select version()")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var version string
	for rows.Next() {
		err := rows.Scan(&version)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("version=%s\n", version)
}
