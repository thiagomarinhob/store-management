package repositories

import (
	"store-management/internal/models"

	"gorm.io/gorm"
)

type StoreRepository interface {
	CreateStore(store *models.Store) error
	FindStoreByID(id string) (*models.Store, error)
	UpdateStore(store *models.Store) error
	DeleteStore(id string) error
	ListStores() ([]models.Store, error)
}

type storeRepository struct {
	db *gorm.DB
}

func NewStoreRepository(db *gorm.DB) StoreRepository {
	return &storeRepository{db: db}
}

func (r *storeRepository) CreateStore(store *models.Store) error {
	return r.db.Create(store).Error
}

func (r *storeRepository) FindStoreByID(id string) (*models.Store, error) {
	var store models.Store
	if err := r.db.First(&store, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &store, nil
}

func (r *storeRepository) UpdateStore(store *models.Store) error {
	return r.db.Save(store).Error
}

func (r *storeRepository) DeleteStore(id string) error {
	return r.db.Delete(&models.Store{}, "id = ?", id).Error
}

func (r *storeRepository) ListStores() ([]models.Store, error) {
	var stores []models.Store
	if err := r.db.Find(&stores).Error; err != nil {
		return nil, err
	}
	return stores, nil
}
