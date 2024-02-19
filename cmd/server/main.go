package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/sidney-cardoso/ecommerce-GO/internal/dto"
	"github.com/sidney-cardoso/ecommerce-GO/internal/entity"
	"github.com/sidney-cardoso/ecommerce-GO/internal/infra/database"
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
	productHandler := NewProductHandler(productDB)

	http.HandleFunc("/products", productHandler.CreateProduct)

	http.ListenAndServe(":8000", nil)
}

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Erro no decode: ", err)
		return
	}
	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Erro no NewProduct: ", err)
		return
	}
	err = h.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println("Erro no create: ", err)
		return
	}
	w.Write([]byte("Deu bom!"))
	w.WriteHeader(http.StatusCreated)
}
