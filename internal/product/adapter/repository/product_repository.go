package repository

import (
	"context"
	"product/internal/product/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// "context"



type ProductRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(client *mongo.Client, dbName, collectionName string) *ProductRepository{
	collection := client.Database(dbName).Collection(collectionName)
	return &ProductRepository{collection: collection}
}

func (repo *ProductRepository) AddProduct(ctx context.Context, product *domain.Product) (*mongo.InsertOneResult, error) {
	return repo.collection.InsertOne(ctx, product)
}

func (repo *ProductRepository) GetProductByID(ctx context.Context, id primitive.ObjectID) (*domain.Product, error) {
    var product domain.Product
    err := repo.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
    if err != nil {
        return nil, err
    }
    return &product, nil
}

func (repo *ProductRepository) UpdateProduct(ctx context.Context, id primitive.ObjectID, update bson.M) (*mongo.UpdateResult, error) {
    return repo.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": update})
}

func (repo *ProductRepository) DeleteProduct(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
    return repo.collection.DeleteOne(ctx, bson.M{"_id": id})
}

func (repo *ProductRepository) GetAllProducts(ctx context.Context) ([]domain.Product, error) {
    var products []domain.Product
    cursor, err := repo.collection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)
    for cursor.Next(ctx) {
        var product domain.Product
        if err := cursor.Decode(&product); err != nil {
            return nil, err
        }
        products = append(products, product)
    }
    return products, nil
}