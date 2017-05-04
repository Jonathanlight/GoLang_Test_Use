package main
import (
  "log"
  "net/http"
)

func serve_files(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path == "/" {
    http.ServeFile(w, r, "index.gohtml")
  }
}

func infoPath(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path == "/info" {
    http.ServeFile(w, r, "index.gohtml")
  }
}

func main() {
  http.HandleFunc("/", serve_files)
  http.HandleFunc("/info", infoPath)
  log.Fatal(http.ListenAndServe(":8000", nil))
}
