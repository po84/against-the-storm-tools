package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type Villager struct {
	Identifier string `json:"id"`
	Name       string `json:"name"`
}

func main() {
	log.Println("opening DB...")
	db, err := sql.Open("sqlite3", "../data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("fetching data...")
	rows, err := db.Query("SELECT identifier, display_name FROM villagers")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var (
		villagers []Villager
		id        string
		name      string
	)

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		villagers = append(villagers, Villager{Identifier: id, Name: name})
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("saving data...")

	filename := "tmp/data.json"
	err = os.MkdirAll(filepath.Dir(filename), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.Encode(villagers)
	log.Println("done")
}
