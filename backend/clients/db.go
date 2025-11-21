package clients

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

func NewDBClient() *sql.DB {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		"user",
		"admin",
		"db",
		5432,
		"database",
	)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("failed to open DB: %v", err)
	}

	// todo: fine tune
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping DB: %v", err)
	}

	log.Println("connected to DB")
	return db
}
