package services

import (
	"store-management/internal/models"
	"store-management/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

type AdminService interface {
	CreateAdmin(admin *models.Admin) error
	GetAdminByID(id string) (*models.Admin, error)
	UpdateAdmin(admin *models.Admin) error
	DeleteAdmin(id string) error
}

type adminService struct {
	adminRepo repositories.AdminRepository
}

func NewAdminService(adminRepo repositories.AdminRepository) AdminService {
	return &adminService{adminRepo: adminRepo}
}

// CreateAdmin implements AdminService.
func (a *adminService) CreateAdmin(admin *models.Admin) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Substitui a senha original pelo hash
	admin.Password = string(hashedPassword)

	// Salva o Admin no banco de dados
	return a.adminRepo.CreateAdmin(admin)
}

// DeleteAdmin implements AdminService.
func (a *adminService) DeleteAdmin(id string) error {
	return a.adminRepo.DeleteAdmin(id)
}

// GetAdminByID implements AdminService.
func (a *adminService) GetAdminByID(id string) (*models.Admin, error) {
	return a.adminRepo.FindAdminByID(id)
}

// UpdateAdmin implements AdminService.
func (a *adminService) UpdateAdmin(admin *models.Admin) error {
	return a.adminRepo.UpdateAdmin(admin)
}
