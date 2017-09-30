package models

import (
	"database/sql"
	"log"
)

type Student struct {
	id       float32
	Name     string
	LastName string
	Score    float32
}

func AllStudents(db *sql.DB) ([]*Student, error) {
	rows, err := db.Query("SELECT * FROM students")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks := make([]*Student, 0)
	for rows.Next() {
		bk := new(Student)
		err := rows.Scan(&bk.id, &bk.Name, &bk.LastName, &bk.Score)
		if err != nil {
			return nil, err
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return bks, nil
}

func AddStudent(db sql.DB, student *Student) {
	insert := `
		INSERT INTO students (Name, LastName, Score)
		VALUES ($1, $2, $3)`

	_, err := db.Exec(insert, student.Name, student.LastName, student.Score)
	if err != nil {
		log.Fatal(err)

	}
}
