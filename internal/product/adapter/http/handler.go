package http

import (
	"context"
	"net/http"
	"product/internal/product/adapter/service"

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
    var req struct {
        ProductName  string `json:"product_name"`
        Stock int    `json:"stock"`
    }
    if err := c.BodyParser(&req); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }
    product, err := h.service.CreateProduct(context.TODO(), req.ProductName, req.Stock)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.Status(http.StatusCreated).JSON(product)
}

func (h *ProductHandler) GetProduct(c *fiber.Ctx) error {
    id, err := primitive.ObjectIDFromHex(c.Params("id"))
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
    }
    product, err := h.service.GetProductByID(context.TODO(), id)
    if err != nil {
        return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(product)
}

func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
    id, err := primitive.ObjectIDFromHex(c.Params("id"))
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
    }
    var req struct {
        ProductName  string `json:"product_name"`
        Stock int    `json:"stock"`
    }
    if err := c.BodyParser(&req); err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }
    product, err := h.service.UpdateProduct(context.TODO(), id, req.ProductName, req.Stock)
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(product)
}

func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
    id, err := primitive.ObjectIDFromHex(c.Params("id"))
    if err != nil {
        return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
    }
    if err := h.service.DeleteProduct(context.TODO(), id); err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.SendStatus(http.StatusNoContent)
}

func (h *ProductHandler) GetAllProducts(c *fiber.Ctx) error {
    products, err := h.service.GetAllProducts(context.TODO())
    if err != nil {
        return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(products)
}