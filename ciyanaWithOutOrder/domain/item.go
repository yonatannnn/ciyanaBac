package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Item struct {
	ID          primitive.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string               `json:"name,omitempty" bson:"name,omitempty"`
	Description string               `json:"description,omitempty" bson:"description,omitempty"`
	Price       float64              `json:"price,omitempty" bson:"price,omitempty"`
	Image       string               `json:"image,omitempty" bson:"image,omitempty"`
	CatagoryID  []primitive.ObjectID `json:"catagoryID,omitempty" bson:"catagoryID,omitempty"`
	Tags        []string             `json:"tags,omitempty" bson:"tags,omitempty"`
	Quantity    int                  `json:"quantity,omitempty" bson:"quantity,omitempty"`
}

type ItemUsecase interface {
	CreateItem(item *Item) error
	GetItem(id string) (*Item, error)
	GetItems() ([]*Item, error)
	UpdateItem(id string, item *Item) error
	DeleteItem(id string) error
	FilterByCategory(id string) ([]*Item, error)
	SearchItem(name string) ([]*Item, error)
	FilterByTag(tag string) ([]*Item, error)
}

type ItemRepository interface {
	CreateItem(item *Item) error
	GetItem(id primitive.ObjectID) (*Item, error)
	GetItems() ([]*Item, error)
	UpdateItem(id primitive.ObjectID, item *Item) error
	DeleteItem(id primitive.ObjectID) error
	FilterByCategory(id primitive.ObjectID) ([]*Item, error)
	SearchItem(name string) ([]*Item, error)
	FilterByTag(tag string) ([]*Item, error)
}
