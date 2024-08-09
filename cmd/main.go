package main

import (
	"log"
	"product/internal/config"
	"product/internal/product/adapter/repository"
	"product/internal/product/adapter/service"
	"product/internal/product/adapter/http"
	"github.com/gofiber/fiber/v2"
)

func main() {
   mongoURI := "mongodb://localhost:27017"

   // Inisialisasi koneksi ke MongoDB
   client := config.InitMongoDB(mongoURI)
   if client == nil {
	   log.Fatal("MongoDB connection failed")
   }

   // Inisialisasi Fiber
	app := fiber.New()

	productRepo := repository.NewProductRepository(client, "productDB", "products")
	productService := service.NewProductService(productRepo)
    productHandler := http.NewProductHandler(productService)

    app.Post("/products", productHandler.CreateProduct)
    app.Get("/products/:id", productHandler.GetProduct)
    app.Put("/products/:id", productHandler.UpdateProduct)
    app.Delete("/products/:id", productHandler.DeleteProduct)
    app.Get("/products", productHandler.GetAllProducts)

   // Jalankan server Fiber
	log.Printf("Successfully connected to MongoDB")
	log.Fatal(app.Listen(":3000"))
}
