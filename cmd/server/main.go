package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/Telekurysh/ozon_test/internal/models"
	"github.com/Telekurysh/ozon_test/pkg/handlers"
)

func main() {
	db := setupDatabase()
	db.AutoMigrate(&models.Post{}, &models.Comment{})

	r := mux.NewRouter()
	r.Handle("/query", handlers.NewGraphQLHandler())
	r.Handle("/playground", handlers.NewPlaygroundHandler())

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Println("Listening on http://localhost:8080")
	log.Fatal(srv.ListenAndServe())
}

func setupDatabase() *gorm.DB {
	dbType := os.Getenv("DB_TYPE")
	var dsn string
	if dbType == "postgres" {
		dsn = os.Getenv("POSTGRES_DSN")
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to connect to postgres database: %v", err)
		}
		return db
	} else {
		db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to connect to sqlite database: %v", err)
		}
		return db
	}
}
