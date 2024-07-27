package controller

import (
	"example/customer-service/model"
	"example/customer-service/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerController interface {
	Signup(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	FindAll(c *gin.Context)
	FindByID(c *gin.Context)
	FindAllAddresses(c *gin.Context)
}

type customerController struct {
	service service.CustomerService
}

func (c *customerController) Signup(ctx *gin.Context) {
	var customer model.Customer
	if err := ctx.ShouldBindBodyWithJSON(&customer); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "failed to bind JSON"})
		return
	}

	if id, err := c.service.Signup(&customer); err == nil {
		ctx.IndentedJSON(http.StatusOK, gin.H{"id": id})
	} else {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to signup"})
	}
}

func (c *customerController) Update(ctx *gin.Context) {
	var customer model.Customer
	if err := ctx.ShouldBindBodyWithJSON(&customer); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "failed to bind JSON"})
		return
	}

	if id, err := strconv.Atoi(ctx.Param("id")); err == nil {
		if err := c.service.Update(uint(id), &customer); err == nil {
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "updated"})
		} else {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to update"})
		}
	}
}

func (c *customerController) Delete(ctx *gin.Context) {
	if id, err := strconv.Atoi(ctx.Param("id")); err == nil {
		if err := c.service.Delete(uint(id)); err == nil {
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "deleted"})
		} else {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to delete"})
		}
	}
}

func (c *customerController) FindAll(ctx *gin.Context) {
	if customers, err := c.service.FindAll(); err == nil {
		ctx.IndentedJSON(http.StatusOK, customers)
	} else {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to find all"})
	}
}

func (c *customerController) FindByID(ctx *gin.Context) {
	if id, err := strconv.Atoi(ctx.Param("id")); err == nil {
		if customer, err := c.service.FindByID(uint(id)); err == nil {
			ctx.IndentedJSON(http.StatusOK, customer)
		} else {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "not found"})
		}
	}
}

func (c *customerController) FindAllAddresses(ctx *gin.Context) {
	if id, err := strconv.Atoi(ctx.Param("id")); err == nil {
		if addresses, err := c.service.FindAllAddresses(uint(id)); err == nil {
			if len(addresses) == 0 {
				ctx.IndentedJSON(http.StatusNoContent, gin.H{"error": "no addresses found"})
				return
			}
			ctx.IndentedJSON(http.StatusOK, addresses)
		} else {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "failed to find addresses"})
		}
	}
}

func NewCustomerController(engine *gin.Engine, service service.CustomerService) {
	c := &customerController{
		service: service,
	}
	api := engine.Group("/api")
	{
		api.GET("/customers", c.FindAll)
		api.POST("/customers", c.Signup)
		api.GET("/customers/:id", c.FindByID)
		api.PUT("/customers/:id", c.Update)
		api.DELETE("/customers/:id", c.Delete)
	}
}
