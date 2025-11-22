package main

import (
	"log"
	"os"
	"path/filepath"
	"sort"

	"github.com/hejare/mega-dashboard-3000/clients"
)

func main() {
	db := clients.NewDBClient()

	// get all files in seed-db directory
	files, err := os.ReadDir("./seeds")
	if err != nil {
		log.Fatalf("failed to read seeds directory: %v", err)
	}

	var sqlFiles []string
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".sql" {
			sqlFiles = append(sqlFiles, file.Name())
		}
	}
	sort.Strings(sqlFiles)

	// iterate over files and execute SQL statements using db
	for _, file := range sqlFiles {
		log.Printf("Applying seed: %s", file)
		content, err := os.ReadFile("./seeds/" + file)
		if err != nil {
			log.Fatalf("failed to read seed file %s: %v", file, err)
		}

		if _, err := db.Exec(string(content)); err != nil {
			log.Fatalf("failed to execute seed file %s: %v", file, err)
		}
	}

	log.Println("Database seeded successfully")
}
