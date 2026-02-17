package main

import (
	"log"
	"net/http"
	"os"
)

func init() {
	err := ConnectToDb()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	defer db.Close()
	http.HandleFunc("/asteroids-scores", handleAsteroidsAPI)

	log.Printf("Server is running on port %s\n", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), nil))
}
