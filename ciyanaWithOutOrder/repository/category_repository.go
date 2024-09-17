package repository

import (
	"ciyana/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type categoryRepository struct {
	dbc *mongo.Collection
	ctx context.Context
}

func NewCategoryRepository(dbc *mongo.Collection, ctx context.Context) domain.CategoryRepository {
	return &categoryRepository{
		dbc: dbc,
		ctx: ctx,
	}
}

func (r *categoryRepository) CreateCategory(category *domain.Category) error {
	_, err := r.dbc.InsertOne(r.ctx, category)
	return err
}

func (r *categoryRepository) GetCategory(id primitive.ObjectID) (*domain.Category, error) {
	var category domain.Category
	err := r.dbc.FindOne(r.ctx, bson.M{"_id": id}).Decode(&category)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) GetCategories() ([]*domain.Category, error) {
	cursor, err := r.dbc.Find(r.ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(r.ctx)

	var categories []*domain.Category
	for cursor.Next(r.ctx) {
		var category domain.Category
		if err := cursor.Decode(&category); err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *categoryRepository) UpdateCategory(id primitive.ObjectID, category *domain.Category) error {
	update := bson.M{}
	if category.Name != "" {
		update["name"] = category.Name
	}
	_, err := r.dbc.UpdateOne(r.ctx, bson.M{"_id": id}, bson.M{"$set": update})
	return err
}

func (r *categoryRepository) DeleteCategory(id primitive.ObjectID) error {
	_, err := r.dbc.DeleteOne(r.ctx, bson.M{"_id": id})
	return err
}
