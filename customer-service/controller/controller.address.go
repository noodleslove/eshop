package controller

import (
	"example/customer-service/model"
	"example/customer-service/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AddressController interface {
	Add(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type addressController struct {
	service service.AddressService
}

// Add adds an address to the database
func (a *addressController) Add(ctx *gin.Context) {
	var address model.Address
	// Bind the JSON body to the address struct
	if err := ctx.ShouldBindBodyWithJSON(&address); err != nil {
		// If there is an error, return a 400 Bad Request response
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Failed to bind JSON"})
		return
	}
	// Add the address to the database
	if rowsAffected, err := a.service.Add(&address); err == nil {
		// If the address was added successfully, return a 200 OK response
		ctx.IndentedJSON(http.StatusOK, gin.H{"RowsAffected": rowsAffected})
	} else {
		// If there was an error adding the address, return a 500 Internal Server Error response
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"Error": "Failed to add address"})
	}
}

// Update updates an address in the database
func (a *addressController) Update(ctx *gin.Context) {
	var address model.Address
	// Bind the JSON body to the address struct
	if err := ctx.ShouldBindBodyWithJSON(&address); err != nil {
		// If there is an error, return a 400 Bad Request response
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Failed to bind JSON"})
		return
	}
	// Update the address in the database
	if id, err := strconv.Atoi(ctx.Param("id")); err == nil {
		// If the ID is valid, update the address
		if err := a.service.Update(uint(id), &address); err == nil {
			// If the address was updated successfully, return a 200 OK response
			ctx.IndentedJSON(http.StatusOK, gin.H{"Message": "Address updated"})
		} else {
			// If there was an error updating the address, return a 500 Internal Server Error response
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"Error": "Failed to update address"})
		}
	} else {
		// If the ID is invalid, return a 400 Bad Request response
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Invalid ID"})
	}
}

// Delete deletes an address from the database
func (a *addressController) Delete(ctx *gin.Context) {
	// Get the ID from the URL parameter
	if id, err := strconv.Atoi(ctx.Param("id")); err == nil {
		// If the ID is valid, delete the address
		if err := a.service.Delete(uint(id)); err == nil {
			// If the address was deleted successfully, return a 200 OK response
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "address deleted"})
		} else {
			// If there was an error deleting the address, return a 500 Internal Server Error response
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"Error": "Failed to delete address"})
		}
	} else {
		// If the ID is invalid, return a 400 Bad Request response
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"Error": "Invalid ID"})
	}
}

// NewAddressController creates a new address controller
func NewAddressController(router *gin.Engine, service service.AddressService) {
	c := &addressController{service}
	api := router.Group("/api")
	{
		api.POST("/address", c.Add)
		api.PUT("/address/:id", c.Update)
		api.DELETE("/address/:id", c.Delete)
	}

}
