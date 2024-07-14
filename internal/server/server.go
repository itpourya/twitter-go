package server

import (
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"twitter-go-api/internal/database"
)

type Server struct {
	port int

	db *gorm.DB
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi("8080")
	NewServer := &Server{
		port: port,

		db: database.New(),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
