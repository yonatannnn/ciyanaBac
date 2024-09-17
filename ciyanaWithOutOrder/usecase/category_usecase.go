package usecase

import (
	"ciyana/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type categoryUsecase struct {
	caregoryRepo domain.CategoryRepository
}

func NewCategoryUsecase(categoryRepo domain.CategoryRepository) domain.CategoryUsecase {
	return &categoryUsecase{
		caregoryRepo: categoryRepo,
	}
}

func (uc *categoryUsecase) CreateCategory(category *domain.Category) error {
	category.ID = primitive.NewObjectID()
	err := uc.caregoryRepo.CreateCategory(category)
	if err != nil {
		return err
	}
	return nil
}

func (uc *categoryUsecase) GetCategory(id string) (*domain.Category, error) {
	ObjId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return uc.caregoryRepo.GetCategory(ObjId)
}

func (uc *categoryUsecase) GetCategories() ([]*domain.Category, error) {
	return uc.caregoryRepo.GetCategories()
}

func (uc *categoryUsecase) UpdateCategory(id string, category *domain.Category) error {
	ObjId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return uc.caregoryRepo.UpdateCategory(ObjId, category)
}

func (uc *categoryUsecase) DeleteCategory(id string) error {
	ObjId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	return uc.caregoryRepo.DeleteCategory(ObjId)
}
