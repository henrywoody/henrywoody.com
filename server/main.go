package main

import (
    "log"
    "os"
    "path"
    "net/http"
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
    http.HandleFunc("/", serveStatic("/index.html", "text/html"))
    http.Handle("/favicon.ico", http.NotFoundHandler())
    http.HandleFunc("/api/asteroids_scores", handleAsteroidsAPI)

    log.Printf("Server is running on port %s\n", os.Getenv("PORT"))
    log.Fatal(http.ListenAndServe(os.Getenv("PORT"), nil))
}

func serveStatic(filename, contentType string) func(http.ResponseWriter, *http.Request) {
    return func(res http.ResponseWriter, req *http.Request) {
        var filepath string
        if filename == "" {
            filepath = path.Join("client/dist/", req.URL.Path)
        } else {
            filepath = path.Join("client/dist/", filename)
        }

        res.Header().Set("Content-Type", contentType)
        http.ServeFile(res, req, filepath)
    }
}