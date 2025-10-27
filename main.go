package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT not found in the enviroment")
	}

	router := chi.NewRouter()

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/health", HandlerRediness)

	v1Router.Get("/error", HandlerError)

	router.Mount("/v1", v1Router)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is running! 🚀"))
	})

	log.Println("Listening on port:", portString)

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
