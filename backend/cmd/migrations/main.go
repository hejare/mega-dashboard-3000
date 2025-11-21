package main

import (
	"fmt"

	"github.com/hejare/mega-dashboard-3000/clients"
)

func main() {
	fmt.Println("running migrations...")
	db := clients.NewDBClient()

	fmt.Println("001 create consultants table")

	tablename := "consultants"

	sql := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
			id SERIAL PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			change_status VARCHAR(100),
			deleted_at TIMESTAMP
		);`, tablename)

	_, err := db.Exec(sql)
	if err != nil {
		fmt.Println("failed to create consultants table")
		panic(err)
	}

	fmt.Println("created consultants table")

	fmt.Println("finished migrations")
}
