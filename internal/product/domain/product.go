package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
    ID			primitive.ObjectID	`bson:"_id,omitempty" json:"product_id,omitempty"`
    ProductName	string				`bson:"product_name" json:"product_name"`
    Stock		int					`bson:"stock" json:"stock"`
}
