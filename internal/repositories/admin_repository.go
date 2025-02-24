package repositories

import (
	"store-management/internal/models"

	"gorm.io/gorm"
)

type AdminRepository interface {
	CreateAdmin(admin *models.Admin) error
	FindAdminByID(id string) (*models.Admin, error)
	UpdateAdmin(admin *models.Admin) error
	DeleteAdmin(id string) error
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepository{db: db}
}

func (r *adminRepository) CreateAdmin(admin *models.Admin) error {
	return r.db.Create(admin).Error
}

func (r *adminRepository) FindAdminByID(id string) (*models.Admin, error) {
	var admin models.Admin
	if err := r.db.First(&admin, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}

func (r *adminRepository) UpdateAdmin(admin *models.Admin) error {
	return r.db.Save(admin).Error
}

func (r *adminRepository) DeleteAdmin(id string) error {
	return r.db.Delete(&models.Admin{}, "id = ?", id).Error
}
