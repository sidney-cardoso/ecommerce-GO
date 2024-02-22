package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/sidney-cardoso/ecommerce-GO/internal/entity"
	"github.com/sidney-cardoso/ecommerce-GO/internal/infra/database"
	"github.com/sidney-cardoso/ecommerce-GO/internal/infra/webserver/handlers"
)

var baseDir string

func init() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to determine file path")
	}
	baseDir = filepath.Dir(filename)
}

func main() {
	err := godotenv.Load(filepath.Join(baseDir, ".env"))
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbDir := filepath.Join(baseDir, "")
	if err := os.MkdirAll(dbDir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create directory: %v", err)
	}

	dbPath := filepath.Join(dbDir, "test.db")
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to initialize database, got error", err)
	}

	db.AutoMigrate(&entity.User{}, &entity.Product{})
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	http.HandleFunc("/products", productHandler.CreateProduct)

	http.ListenAndServe(":8000", nil)
}
