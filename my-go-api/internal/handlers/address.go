package handlers

import (
	"net/http"
	"strconv"

	"my-go-api/internal/models"
	"my-go-api/internal/repository" // Import the repository package

	"github.com/gin-gonic/gin"
)

type AddressHandler struct {
	Repo repository.AddressRepository // Inject the interface, NOT *sql.DB
}

// Update the constructor to accept the repository interface
func NewAddressHandler(repo repository.AddressRepository) *AddressHandler {
	return &AddressHandler{Repo: repo}
}

// GET /api/v1/address
func (h *AddressHandler) GetAddresses(c *gin.Context) {
	userID := c.MustGet("userID").(uint)

	// Call the repository layer instead of h.DB.Query
	addresses, err := h.Repo.GetByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch addresses"})
		return
	}
	c.JSON(http.StatusOK, addresses)
}

// POST /api/v1/address
func (h *AddressHandler) CreateAddress(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	var input models.Address

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.UserID = userID

	// Call the repository layer instead of h.DB.QueryRow
	if err := h.Repo.Create(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save address details"})
		return
	}
	c.JSON(http.StatusCreated, input)
}

// PUT /api/v1/address?id=1
func (h *AddressHandler) UpdateAddress(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	addressIDStr := c.Query("id")

	addressID, err := strconv.Atoi(addressIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Valid Address ID query parameter is required"})
		return
	}

	var input models.Address
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call h.Repo.Update instead of h.DB.Exec
	rowsAffected, err := h.Repo.Update(&input, uint(addressID), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update address"})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Address record not found or access denied"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Address updated successfully"})
}

// DELETE /api/v1/address?id=1
func (h *AddressHandler) DeleteAddress(c *gin.Context) {
	userID := c.MustGet("userID").(uint)
	addressIDStr := c.Query("id")

	addressID, err := strconv.Atoi(addressIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Valid Address ID query parameter is required"})
		return
	}

	// Call h.Repo.Delete instead of h.DB.Exec
	rowsAffected, err := h.Repo.Delete(uint(addressID), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete address"})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Address record not found or access denied"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Address deleted successfully"})
}
