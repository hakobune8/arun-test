package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func staticDir() string {
	for _, dir := range []string{"client", filepath.Join("..", "client")} {
		if _, err := os.Stat(filepath.Join(dir, "index.html")); err == nil {
			return dir
		}
	}
	return "client"
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	dir := staticDir()
	if r.URL.Path != "/" {
		if assetPath, ok := staticAssetPath(r.URL.Path); ok {
			http.ServeFile(w, r, assetPath)
			return
		}
		http.NotFound(w, r)
		return
	}
	if _, err := os.Stat(filepath.Join(dir, "index.html")); err == nil {
		http.ServeFile(w, r, filepath.Join(dir, "index.html"))
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("arun-test service\n"))
}

func staticAssetPath(urlPath string) (string, bool) {
	dir := staticDir()
	clean := path.Clean("/" + urlPath)
	switch {
	case clean == "/styles.css":
		return filepath.Join(dir, "styles.css"), true
	case strings.HasPrefix(clean, "/src/") && strings.HasSuffix(clean, ".js"):
		return filepath.Join(dir, strings.TrimPrefix(clean, "/")), true
	default:
		return "", false
	}
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/healthz", healthzHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
