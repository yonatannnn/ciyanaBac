package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	ParentId string             `json:"parent_id" bson:"parent_id"`
}

type CategoryUsecase interface {
	CreateCategory(category *Category) error
	GetCategory(id string) (*Category, error)
	GetCategories() ([]*Category, error)
	UpdateCategory(id string, category *Category) error
	DeleteCategory(id string) error
}

type CategoryRepository interface {
	CreateCategory(category *Category) error
	GetCategory(id primitive.ObjectID) (*Category, error)
	GetCategories() ([]*Category, error)
	UpdateCategory(id primitive.ObjectID, category *Category) error
	DeleteCategory(id primitive.ObjectID) error
}
