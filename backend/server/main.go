package main

import (
	"log"
	"net/http"
	"os"
	"path"
)

func init() {
	err := ConnectToDb()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	defer db.Close()
	http.HandleFunc("/css/", serveStatic("", "text/css"))
	http.HandleFunc("/js/", serveStatic("", "application/javascript"))
	http.HandleFunc("/icon/", serveStatic("", "image/png"))
	http.HandleFunc("/favicon.ico", serveStatic("/favicon.ico", "image/ico"))
	http.HandleFunc("/manifest.json", serveStatic("/manifest.json", "application/json"))
	http.HandleFunc("/", serveStatic("/index.html", "text/html"))
	http.HandleFunc("/api/asteroids_scores", handleAsteroidsAPI)

	log.Printf("Server is running on port %s\n", os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), nil))
}

func serveStatic(filename, contentType string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var filepath string
		if filename == "" {
			filepath = path.Join("client/dist/", r.URL.Path)
		} else {
			filepath = path.Join("client/dist/", filename)
		}

		w.Header().Set("Content-Type", contentType)
		http.ServeFile(w, r, filepath)
	}
}
