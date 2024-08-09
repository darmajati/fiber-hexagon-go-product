package service

import (
	"context"
	"product/internal/product/adapter/repository"
	"product/internal/product/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(ctx context.Context, name string, stock int) (*domain.Product, error) {
    product := &domain.Product{
        ID:    primitive.NewObjectID(),
        ProductName:  name,
        Stock: stock,
    }
    _, err := s.repo.AddProduct(ctx, product)
    if err != nil {
        return nil, err
    }
    return product, nil
}

func (s *ProductService) GetProductByID(ctx context.Context, id primitive.ObjectID) (*domain.Product, error) {
    return s.repo.GetProductByID(ctx, id)
}

func (s *ProductService) UpdateProduct(ctx context.Context, id primitive.ObjectID, name string, stock int) (*domain.Product, error) {
    update := bson.M{"product_name": name, "stock": stock}
    _, err := s.repo.UpdateProduct(ctx, id, update)
    if err != nil {
        return nil, err
    }
    return s.GetProductByID(ctx, id)
}

func (s *ProductService) DeleteProduct(ctx context.Context, id primitive.ObjectID) error {
    _, err := s.repo.DeleteProduct(ctx, id)
    return err
}

func (s *ProductService) GetAllProducts(ctx context.Context) ([]domain.Product, error) {
    return s.repo.GetAllProducts(ctx)
}