package services

import (
	"store-management/internal/models"
	"store-management/internal/repositories"
)

type StoreService interface {
	CreateStore(store *models.Store) error
	GetStoreByID(id string) (*models.Store, error)
	UpdateStore(store *models.Store) error
	DeleteStore(id string) error
	ListStores() ([]models.Store, error)
}

type storeService struct {
	storeRepo repositories.StoreRepository
}

func NewStoreService(storeRepo repositories.StoreRepository) StoreService {
	return &storeService{storeRepo: storeRepo}
}

func (s *storeService) CreateStore(store *models.Store) error {
	return s.storeRepo.CreateStore(store)
}

func (s *storeService) GetStoreByID(id string) (*models.Store, error) {
	return s.storeRepo.FindStoreByID(id)
}

func (s *storeService) UpdateStore(store *models.Store) error {
	return s.storeRepo.UpdateStore(store)
}

func (s *storeService) DeleteStore(id string) error {
	return s.storeRepo.DeleteStore(id)
}

func (s *storeService) ListStores() ([]models.Store, error) {
	return s.storeRepo.ListStores()
}
