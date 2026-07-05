package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

const staticDir = "client"

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		if assetPath, ok := staticAssetPath(r.URL.Path); ok {
			http.ServeFile(w, r, assetPath)
			return
		}
		http.NotFound(w, r)
		return
	}
	if _, err := os.Stat(path.Join(staticDir, "index.html")); err == nil {
		http.ServeFile(w, r, path.Join(staticDir, "index.html"))
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("arun-test service\n"))
}

func staticAssetPath(urlPath string) (string, bool) {
	clean := path.Clean("/" + urlPath)
	switch {
	case clean == "/styles.css":
		return path.Join(staticDir, "styles.css"), true
	case strings.HasPrefix(clean, "/src/") && strings.HasSuffix(clean, ".js"):
		return path.Join(staticDir, strings.TrimPrefix(clean, "/")), true
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
	log.Fatal(http.ListenAndServe(":8080", mux))
}
