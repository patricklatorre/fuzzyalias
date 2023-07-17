package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Config struct {
	LinkMap        map[string]string `json:"links"`
	TimeoutSeconds int               `json:"timeoutSeconds"`
	ThrottleLimit  int               `json:"throttleLimit"`
}

var config Config

func main() {
	var (
		configFlag = flag.String("config", "", "config file path")
		localFlag  = flag.Bool("local", false, "sets host to 127.0.0.1")
	)

	flag.Parse()
	mustLoadConfig(*configFlag)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Throttle(config.ThrottleLimit))
	r.Use(middleware.Timeout(time.Second * time.Duration(config.TimeoutSeconds)))
	r.Use(middleware.Recoverer)
	r.Use(middleware.NoCache)
	r.Get("/*", fuzzyRedirect)

	addr := getHost(*localFlag) + getPort()
	err := http.ListenAndServe(addr, r)

	if err != http.ErrServerClosed {
		log.Panicln(err)
	}
}

func mustLoadConfig(path string) {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		log.Panicln(err)
	}

	err = json.Unmarshal(fileContent, &config)
	if err != nil {
		log.Panicln(err)
	}

	log.Println("Loaded config:", path)
}

func fuzzyRedirect(w http.ResponseWriter, r *http.Request) {
	link, exists := searchNearestLink(r.URL.String())
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`<p style="padding: 1rem; font-family: Arial, sans-serif;">No matches found.</p>`))
		return
	}

	http.Redirect(w, r, link, http.StatusTemporaryRedirect)
}

func getHost(local bool) string {
	if local {
		return "127.0.0.1"
	}

	return ""
}

func getPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		return ":8090"
	}

	return ":" + port
}
