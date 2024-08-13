package http

import (
	"context"
	"fmt"
	"net/http"
	"product/internal/product/adapter/service"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
    return &ProductHandler{service: service}
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
    start := time.Now()

    var req struct {
        ProductName  string `json:"product_name"`
        Stock int    `json:"stock"`
    }
    if err := c.BodyParser(&req); err != nil {
        return c.Status(http.StatusBadRequest).JSON(struct {
            StatusCode int    `json:"status_code"`
            Message    string `json:"message"`
        }{
            StatusCode: fiber.StatusBadRequest,
            Message:    err.Error(),
        })
    }
    product, err := h.service.CreateProduct(context.TODO(), req.ProductName, req.Stock)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(struct {
            StatusCode int    `json:"status_code"`
            Message    string `json:"message"`
        }{
            StatusCode: fiber.StatusInternalServerError,
            Message:    err.Error(),
        })
    }
    duration:= time.Since(start)
    fmt.Println("Create Product took: ", duration)

    return c.Status(http.StatusCreated).JSON(struct {
        StatusCode int         `json:"status_code"`
        Message    string      `json:"message"`
        Data       interface{} `json:"data"`
    }{
        StatusCode: fiber.StatusCreated,
        Message:    "Data berhasil dibuat",
        Data:       product,
    })
}

func (h *ProductHandler) GetProduct(c *fiber.Ctx) error {
    start := time.Now()

    id, err := primitive.ObjectIDFromHex(c.Params("id"))
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(struct {
            StatusCode int    `json:"status_code"`
            Message    string `json:"message"`
        }{
            StatusCode: fiber.StatusBadRequest,
            Message:    "Invalid ID",
        })
    }
    product, err := h.service.GetProductByID(context.TODO(), id)
    if err != nil {
        return c.Status(http.StatusNotFound).JSON(struct {
            StatusCode int    `json:"status_code"`
            Message    string `json:"message"`
        }{
            StatusCode: fiber.StatusNotFound,
            Message:    "Data tidak ditemukan",
        })
    }

    duration:= time.Since(start)
    fmt.Println("Get Product By Id took: ", duration)

    return c.Status(http.StatusOK).JSON(struct {
        StatusCode int         `json:"status_code"`
        Message    string      `json:"message"`
        Data       interface{} `json:"data"`
    }{
        StatusCode: fiber.StatusOK,
        Message:    "Data berhasil ditemukan",
        Data:       product,
    })
}

func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
    start := time.Now()

    id, err := primitive.ObjectIDFromHex(c.Params("id"))
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(struct {
            StatusCode  int     `json:"status_code"`
            Message     string  `json:"message"`
        }{
            StatusCode: fiber.StatusBadRequest,
            Message: "Invalid ID",
        })
    }
    var req struct {
        ProductName  string `json:"product_name"`
        Stock int    `json:"stock"`
    }
    if err := c.BodyParser(&req); err != nil {
        return c.Status(http.StatusBadRequest).JSON(struct {
            StatusCode int    `json:"status_code"`
            Message    string `json:"message"`
        }{
            StatusCode: fiber.StatusBadRequest,
            Message:    err.Error(),
        })
    }
    product, err := h.service.UpdateProduct(context.TODO(), id, req.ProductName, req.Stock)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(struct {
            StatusCode int    `json:"status_code"`
            Message    string `json:"message"`
        }{
            StatusCode: fiber.StatusInternalServerError,
            Message:    err.Error(),
        })
    }

    duration:= time.Since(start)
    fmt.Println("Update Product took: ", duration)

    return c.Status(http.StatusCreated).JSON(struct {
        StatusCode int         `json:"status_code"`
        Message    string      `json:"message"`
        Data       interface{} `json:"data"`
    }{
        StatusCode: fiber.StatusCreated,
        Message:    "Data berhasil diedit",
        Data:       product,
    })
}

func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
    start := time.Now()

    id, err := primitive.ObjectIDFromHex(c.Params("id"))
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(struct {
            StatusCode  int     `json:"status_code"`
            Message     string  `json:"message"`
        }{
            StatusCode: fiber.StatusBadRequest,
            Message: "Invalid ID",
        })
    }
    if err := h.service.DeleteProduct(context.TODO(), id); err != nil {
        return c.Status(http.StatusInternalServerError).JSON(struct {
            StatusCode int    `json:"status_code"`
            Message    string `json:"message"`
        }{
            StatusCode: fiber.StatusInternalServerError,
            Message:    err.Error(),
        })
    }

    duration:= time.Since(start)
    fmt.Println("Delete Product took: ", duration)

    return c.Status(http.StatusOK).JSON(struct {
        StatusCode int    `json:"status_code"`
        Message    string `json:"message"`
    }{
        StatusCode: fiber.StatusOK,
        Message:    "Data berhasil dihapus",
    })
}

func (h *ProductHandler) GetAllProducts(c *fiber.Ctx) error {
    start := time.Now()

    products, err := h.service.GetAllProducts(context.TODO())
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(struct {
            StatusCode int    `json:"status_code"`
            Message    string `json:"message"`
        }{
            StatusCode: fiber.StatusInternalServerError,
            Message:    err.Error(),
        })
    }

    duration:= time.Since(start)
    fmt.Println("Get All Products took: ", duration)

    return c.Status(http.StatusOK).JSON(struct {
        StatusCode int         `json:"status_code"`
        Message    string      `json:"message"`
        Data       interface{} `json:"data"`
    }{
        StatusCode: fiber.StatusOK,
        Message:    "Data berhasil ditemukan",
        Data:       products,
    })
}