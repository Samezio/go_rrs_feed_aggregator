package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/samezio/rrs_aggregator/internal/database"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")
	port := 8080
	if portString := os.Getenv("PORT"); portString == "" {
		log.Fatal("'PORT' env isn't present")
	} else if p, err := strconv.Atoi(portString); err != nil {
		log.Fatal("'PORT' in env isn't number", err)
	} else {
		port = p
	}

	apiCfg := apiConfig{
		DB: nil,
	}

	if dbUrl := os.Getenv("DB_URL"); dbUrl == "" {
		log.Fatal("'DB_URL' env isn't present")
	} else if connection, err := sql.Open("postgres", dbUrl); err != nil {
		log.Fatal("Error in creating DB connection", err)
	} else {
		apiCfg.DB = database.New(connection)
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "OPTIONS", "DELETE", "PATCH"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handleReadiness) //Health check route
	v1Router.Get("/error", handlerError)
	v1Router.Post("/users", apiCfg.handleCreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handleGetUser))
	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%d", port),
	}

	log.Printf("Server starting at %s\n", fmt.Sprintf(":%d", port))
	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
