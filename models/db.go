package models

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func NewDB(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func CreateSchema(db *sql.DB) {
	schema := `
        CREATE TABLE IF NOT EXISTS students (
          id SERIAL PRIMARY KEY,
          Name TEXT,
          LastName TEXT,
          Score DECIMAL
        );
    `
	execDb(db, schema)
}

func PopulateDatabase(db *sql.DB) {
	insert := `
		INSERT INTO students (Name, LastName, Score)
		VALUES ('Artem', 'Karp', 10)`
	execDb(db, insert)
}

func execDb(db *sql.DB, schema string) {
	_, err := db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}
}