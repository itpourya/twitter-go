package database

import (
	"gorm.io/gorm"
	"log"
	"twitter-go-api/internal/entity"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/sqlite"
)

func New() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}

	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		panic("Failed : Unable to migrate your postgres database")
	}

	return db
}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func Close() {
	db := New()
	database, _ := db.DB()
	err := database.Close()
	if err != nil {
		log.Fatal(err)
	}
}
