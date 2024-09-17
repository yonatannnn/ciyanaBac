package usecase

import (
	"ciyana/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type itemUsecase struct {
	itemRepo domain.ItemRepository
}

func NewItemUsecase(itemRepo domain.ItemRepository) domain.ItemUsecase {
	return &itemUsecase{
		itemRepo: itemRepo,
	}
}

func (uc *itemUsecase) CreateItem(item *domain.Item) error {
	item.ID = primitive.NewObjectID()
	return uc.itemRepo.CreateItem(item)
}

func (uc *itemUsecase) GetItem(id string) (*domain.Item, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	item, err := uc.itemRepo.GetItem(objID)
	if err != nil {
		return nil, err
	}
	return item, nil

}

func (uc *itemUsecase) GetItems() ([]*domain.Item, error) {
	items, err := uc.itemRepo.GetItems()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (uc *itemUsecase) UpdateItem(id string, item *domain.Item) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	err = uc.itemRepo.UpdateItem(objID, item)
	if err != nil {
		return err
	}
	return nil

}

func (uc *itemUsecase) DeleteItem(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	err = uc.itemRepo.DeleteItem(objID)
	if err != nil {
		return err
	}
	return nil

}

func (uc *itemUsecase) FilterByCategory(id string) ([]*domain.Item, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return uc.itemRepo.FilterByCategory(objID)
}

func (uc *itemUsecase) SearchItem(name string) ([]*domain.Item, error) {
	return uc.itemRepo.SearchItem(name)
}

func (uc *itemUsecase) FilterByTag(tag string) ([]*domain.Item, error) {
	return uc.itemRepo.FilterByTag(tag)
}
