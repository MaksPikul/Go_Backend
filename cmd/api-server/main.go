package main

import (
	"fmt"
	"go_backend/internal/config"
	bucket "go_backend/internal/data/Bucket"
	cache "go_backend/internal/data/Cache"
	repositories "go_backend/internal/data/Repositories"
	"log/slog"
	"os"

	"go_backend/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		// panic before because slog not initialized
		panic(fmt.Errorf("error loading .env file: %w", err))
	}

	logFile := logger.InitLogger()
	defer logFile.Close()

	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("fatal error: failed to load config", slog.Any("error", err))
		os.Exit(1)
	}

	bucketUploader, err := bucket.GetBucketUploader(cfg.Cloud)
	if err != nil {
		slog.Error("fatal error: failed to initialise bucket loader", slog.Any("error", err))
		os.Exit(1)
	}

	// Maybe Set Var that chooses wether to cache or not
	cache, err := cache.GetCacheClient(cfg.Cloud)
	if err != nil {
		slog.Error("fatal error: failed to initialise cache client", slog.Any("error", err))
		os.Exit(1)
	}

	// Create Decorator

	repos := repositories.NewRepositories(bucketUploader)

	// Wrap repo with decorator

	// Start Handlers

	// Start Server
}

func homePage(c *gin.Context) {
	c.String(http.StatusOK, "This is homepage")
}

/*
sample loggin middleware

func LoggerMiddleware(c *gin.Context) {
    path := c.Request.URL.Path
    method := c.Request.Method
    // Before handler
    log.Printf("Started %s %s", method, path)

    c.Next() // call next middleware or handler

    // After handler
    status := c.Writer.Status()
    log.Printf("Completed %d %s", status, path)
}
*/

/*
srv := server.NewServer(cfg)
    if err := srv.Run(); err != nil {
        log.Fatal(err)
    }

type Server struct {
	router *gin.Engine
	cfg    *config.Config
}

func NewServer(cfg *config.Config) *Server {
	r := gin.Default()

	s := &Server{
		router: r,
		cfg:    cfg,
	}

	s.routes()
	return s
}

func (s *Server) routes() {
	s.router.GET("/health", handlers.HealthCheck)
	s.router.GET("/users/:id", handlers.GetUser)
	// Add more routes here
}

func (s *Server) Run() error {
	return s.router.Run(":" + s.cfg.Port)
}
*/

// might not need, might need in User Service
//DB, err := rds.GetDB(cfg.DB)
