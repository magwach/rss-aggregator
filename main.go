package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	// "time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/magwach/rss-aggregator/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT not found in the enviroment")
	}

	dbUrl := os.Getenv("DB_URL")

	if dbUrl == "" {
		log.Fatal("DB URL not found in the enviroment")
	}

	conn, err := sql.Open("postgres", dbUrl)

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	queries := database.New(conn)

	apiCfg := apiConfig{
		DB: queries,
	}

	// go scrape(queries, 10, time.Minute)

	router := chi.NewRouter()
	v1Router := chi.NewRouter()

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

	v1Router.Get("/health", HandlerRediness)

	v1Router.Get("/error", HandlerError)

	v1Router.Post("/user/create", apiCfg.HandlerCreateUser)

	v1Router.Get("/user/get", apiCfg.GetUserMiddleware(apiCfg.HandlerGetUser))

	v1Router.Post("/feed/create", apiCfg.GetUserMiddleware(apiCfg.HandleCreateFeed))

	v1Router.Get("/feed/get", apiCfg.GetAllFeeds)

	v1Router.Post("/feed/follow", apiCfg.GetUserMiddleware(apiCfg.HandleFollowFeed))

	v1Router.Get("/feed/follow", apiCfg.GetUserMiddleware(apiCfg.HandleGetAllFollowingFeeds))

	v1Router.Post("/feed/get", apiCfg.GetUserMiddleware(apiCfg.HandleFollowFeed))

	v1Router.Delete("/feed/unfollow/{feedId}", apiCfg.GetUserMiddleware(apiCfg.HandleUnfollowFeed))

	v1Router.Get("/posts/get", apiCfg.GetUserMiddleware(apiCfg.HandleGetPostsFromFollowing))

	router.Mount("/v1", v1Router)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is running! 🚀"))
	})

	log.Println("Listening on port:", portString)

	err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
