package handlers

import (
	"net/http"
	"store-management/internal/models"
	"store-management/internal/services"

	"github.com/gin-gonic/gin"
)

type StoreHandler struct {
	storeService services.StoreService
}

func NewStoreHandler(storeService services.StoreService) *StoreHandler {
	return &StoreHandler{storeService: storeService}
}

func (h *StoreHandler) CreateStore(c *gin.Context) {
	var store models.Store
	if err := c.ShouldBindJSON(&store); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := h.storeService.CreateStore(&store); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create store"})
		return
	}

	c.JSON(http.StatusCreated, store)
}

func (h *StoreHandler) GetStoreByID(c *gin.Context) {
	id := c.Param("id")

	store, err := h.storeService.GetStoreByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Store not found"})
		return
	}

	c.JSON(http.StatusOK, store)
}

func (h *StoreHandler) UpdateStore(c *gin.Context) {
	id := c.Param("id")

	var store models.Store
	if err := c.ShouldBindJSON(&store); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	store.ID = id
	if err := h.storeService.UpdateStore(&store); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update store"})
		return
	}

	c.JSON(http.StatusOK, store)
}

func (h *StoreHandler) DeleteStore(c *gin.Context) {
	id := c.Param("id")

	if err := h.storeService.DeleteStore(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete store"})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *StoreHandler) ListStores(c *gin.Context) {
	stores, err := h.storeService.ListStores()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list stores"})
		return
	}

	c.JSON(http.StatusOK, stores)
}
