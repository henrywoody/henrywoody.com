package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func handleAsteroidsAPI(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	switch r.Method {
	case "GET":
		getAsteroidsScores(w, r)
	case "POST":
		postAsteroidsScore(w, r)
	}
}

type AsteroidScore struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func getAsteroidsScores(w http.ResponseWriter, r *http.Request) {
	queryText := "SELECT name, score FROM asteroids_score ORDER BY score DESC"
	var queryValues []interface{}

	if limit := r.URL.Query().Get("limit"); limit != "" {
		queryText += " LIMIT $1"
		queryValues = append(queryValues, limit)
	}

	rows, err := db.Query(queryText, queryValues...)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	records := []AsteroidScore{}
	for rows.Next() {
		var record AsteroidScore
		if err := rows.Scan(&record.Name, &record.Score); err != nil {
			log.Fatal(err)
		}

		records = append(records, record)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(records)
	if err != nil {
		log.Fatal(err)
	}
}

func postAsteroidsScore(w http.ResponseWriter, r *http.Request) {
	var record AsteroidScore
	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		log.Fatal(err)
	}

	if len(record.Name) > 3 {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		msg := `{"error": "Name cannot exceed 3 characters"}`
		w.Write([]byte(msg))
		return
	}

	queryText := "INSERT INTO asteroids_score (name, score) VALUES ($1, $2)"
	queryValues := []interface{}{strings.ToUpper(record.Name), record.Score}

	_, err = db.Exec(queryText, queryValues...)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(record)
	if err != nil {
		log.Fatal(err)
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type")
}
