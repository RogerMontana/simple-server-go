package main

import (
	"./models"
	"database/sql"
	"log"
	"net/http"
	"encoding/json"
	"fmt"
)

type Env struct {
	db *sql.DB
}

type Student struct {
	Name     string
	LastName string
	Score    float32
}

func main() {
	db, err := models.NewDB("postgres://pguser:pguser@localhost:5432/pgdb?sslmode=disable")
	if err != nil {
		log.Panic(err)
	}
	env := &Env{db: db}
	models.CreateSchema(db)
	models.PopulateDatabase(db)
	http.Handle("/students", studentsIndex(env))
	http.Handle("/student", studentAdd(env))
	http.ListenAndServe(":3000", nil)
}

func studentsIndex(env *Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, http.StatusText(405), 405)
			return
		}
		bks, err := models.AllStudents(env.db)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			log.Print(err)
			return
		}
		b, _ := json.Marshal(bks)
		w.Write([]byte(b))
	})
}

func studentAdd(env *Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var s Student
		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&s)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		fmt.Println(s.Name)
	})
}