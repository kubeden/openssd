package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/kubeden/openssd/types"

	"github.com/adrg/frontmatter"
	"github.com/gorilla/mux"
	"github.com/russross/blackfriday/v2"
)

type Article struct {
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Slug    string    `json:"slug"`
	Date    time.Time `json:"date"`
	Excerpt string    `json:"excerpt"`
}

var articles []Article
var config types.Config

func StartServer(cfg types.Config) error {
	config = cfg
	loadArticles()

	r := mux.NewRouter()

	r.Use(corsMiddleware)

	r.HandleFunc("/api/articles", getArticles).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/articles/{slug}", getArticle).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/readme", getGithubReadme).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/info", getGithubInfo).Methods("GET", "OPTIONS")

	log.Println("API server starting on :8081")
	return http.ListenAndServe(":8081", r)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, HX-Request, HX-Trigger, HX-Target, HX-Current-URL")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func loadArticles() {
	files, err := ioutil.ReadDir("./articles")
	if err != nil {
		log.Fatalf("Error reading articles directory: %v", err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".md" {
			article, err := parseArticle(file.Name())
			if err != nil {
				log.Printf("Error parsing article %s: %v", file.Name(), err)
				continue
			}
			articles = append(articles, article)
			log.Printf("Loaded article: %s", article.Title)
		}
	}

	sort.Slice(articles, func(i, j int) bool {
		return articles[i].Date.After(articles[j].Date)
	})

	log.Printf("Loaded %d articles", len(articles))
}

func parseArticle(filename string) (Article, error) {
	content, err := ioutil.ReadFile(filepath.Join("./articles", filename))
	if err != nil {
		return Article{}, fmt.Errorf("error reading file %s: %v", filename, err)
	}

	var article Article
	rest, err := frontmatter.Parse(strings.NewReader(string(content)), &article)
	if err != nil {
		return Article{}, fmt.Errorf("error parsing frontmatter for %s: %v", filename, err)
	}

	article.Content = string(blackfriday.Run(rest))

	if article.Slug == "" {
		article.Slug = strings.TrimSuffix(filename, filepath.Ext(filename))
	}

	return article, nil
}

func getArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(articles); err != nil {
		log.Printf("Error encoding articles: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	log.Printf("Sent %d articles", len(articles))
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]

	for _, article := range articles {
		if article.Slug == slug {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(article)
			return
		}
	}

	http.NotFound(w, r)
}

func getGithubReadme(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/main/%s", config.GithubUsername, config.GithubRepo, config.ReadmeFile)
	fetchAndServeMarkdown(w, r, url)
}

func getGithubInfo(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/main/%s", config.GithubUsername, config.GithubRepo, config.InfoFile)
	fetchAndServeMarkdown(w, r, url)
}

func fetchAndServeMarkdown(w http.ResponseWriter, r *http.Request, url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching markdown: %v", err)
		http.Error(w, "Failed to fetch markdown", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading markdown: %v", err)
		http.Error(w, "Failed to read markdown", http.StatusInternalServerError)
		return
	}

	// Convert markdown to HTML
	html := blackfriday.Run(body)

	w.Header().Set("Content-Type", "text/html")
	w.Write(html)
}
