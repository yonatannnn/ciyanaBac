package repository

import (
	"ciyana/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type itemRepository struct {
	dbc *mongo.Collection
	ctx context.Context
}

func NewItemRepository(dbc *mongo.Collection, ctx context.Context) domain.ItemRepository {
	return &itemRepository{
		dbc: dbc,
		ctx: ctx,
	}
}

func (r *itemRepository) CreateItem(item *domain.Item) error {
	_, err := r.dbc.InsertOne(r.ctx, item)
	return err
}

func (r *itemRepository) GetItem(id primitive.ObjectID) (*domain.Item, error) {
	var item domain.Item
	err := r.dbc.FindOne(r.ctx, bson.M{"_id": id}).Decode(&item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *itemRepository) GetItems() ([]*domain.Item, error) {
	cursor, err := r.dbc.Find(r.ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(r.ctx)

	var items []*domain.Item
	for cursor.Next(r.ctx) {
		var item domain.Item
		if err := cursor.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, &item)
	}
	return items, nil
}

func (r *itemRepository) UpdateItem(id primitive.ObjectID, item *domain.Item) error {
	update := bson.M{}
	if item.Name != "" {
		update["name"] = item.Name
	}
	if item.Description != "" {
		update["description"] = item.Description
	}
	if item.Price != 0 {
		update["price"] = item.Price
	}
	if item.Quantity != 0 {
		update["quantity"] = bson.M{"$inc": item.Quantity}
	}
	if item.Tags != nil {
		update["tags"] = bson.M{"$each": item.Tags}
	}
	if len(update) == 0 {
		return nil
	}
	_, err := r.dbc.UpdateOne(r.ctx, bson.M{"_id": id}, bson.M{"$set": update})
	return err
}

func (r *itemRepository) DeleteItem(id primitive.ObjectID) error {
	_, err := r.dbc.DeleteOne(r.ctx, bson.M{"_id": id})
	return err
}

func (r *itemRepository) FilterByCategory(id primitive.ObjectID) ([]*domain.Item, error) {
	cursor, err := r.dbc.Find(r.ctx, bson.M{"category": id})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(r.ctx)

	var items []*domain.Item
	for cursor.Next(r.ctx) {
		var item domain.Item
		if err := cursor.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, &item)
	}
	return items, nil
}

func (r *itemRepository) SearchItem(name string) ([]*domain.Item, error) {
	cursor, err := r.dbc.Find(r.ctx, bson.M{"name": bson.M{"$regex": name, "$options": "i"}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(r.ctx)

	var items []*domain.Item
	for cursor.Next(r.ctx) {
		var item domain.Item
		if err := cursor.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, &item)
	}
	return items, nil
}

func (r *itemRepository) FilterByTag(tag string) ([]*domain.Item, error) {
	cursor, err := r.dbc.Find(r.ctx, bson.M{"tags": tag})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(r.ctx)

	var items []*domain.Item
	for cursor.Next(r.ctx) {
		var item domain.Item
		if err := cursor.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, &item)
	}
	return items, nil
}
