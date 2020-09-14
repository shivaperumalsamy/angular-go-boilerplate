package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gobuffalo/envy"
	"github.com/rs/cors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	chandler "github.com/sivagasc/angular-go-boilerplate/pkg/handlers"
)

//Common const
const (
	EnvFile       string = "env"
	APIKey        string = "apikey"
	APIKeyDoc     string = "API key"
	APIPathPrefix string = "/api/v1"
)

// Article ...
type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Articles ...
type Articles []Article

func allArticles() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		articles := Articles{
			Article{Title: "Test Title", Desc: "Test description", Content: "Sample content"},
			Article{Title: "Title 2", Desc: "Test description2", Content: "Sample content2"},
			Article{Title: "Test Title 3", Desc: "Test description3", Content: "Sample content3"},
		}
		fmt.Println("Endpoint hit: All Articles endpoint")
		json.NewEncoder(w).Encode(articles)
	})
}

func sample() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
}

func main() {
	// handler()
	logger := &log.Logger
	environment := envy.Get("ENVIRONMENT", "dev")

	if environment == "prod" {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		logger.Info().Msg("*** Production Configuration ***")
		logger.Debug().Msg("*** Debug Logging Disabled ***") // Won't display
	} else {
		zl := logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		logger = &zl
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		logger.Info().Msg("*** Non-production Configuration ***")
		logger.Debug().Msg("*** Debug Logging Enabled ***")
	}

	r := mux.NewRouter()
	r.Handle("/", chandler.Hello())

	s := r.PathPrefix(APIPathPrefix).Subrouter()
	s.Handle("/", sample())
	s.Handle("/article", allArticles()).Methods("GET")

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	// log.Info().Msg("Port:" + os.Getenv("PORT"))
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowCredentials: true,
	})

	handler := corsHandler.Handler(loggedRouter)
	logger.Fatal().Msg(http.ListenAndServe(":4201", handler).Error())
}
