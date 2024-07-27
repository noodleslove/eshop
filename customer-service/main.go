package main

import (
	"example/customer-service/controller"
	"example/customer-service/model"
	"example/customer-service/repository"
	"example/customer-service/service"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// initDB is a function to initialize the database connection
func initDB() *gorm.DB {
	// Get the connection string from the environment
	dsn := os.Getenv("CONNECTION_STRING")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	// Migrate the database
	if err = db.AutoMigrate(&model.Customer{}); err != nil {
		panic("Failed to migrate Customer model")
	}
	if err = db.AutoMigrate(&model.Address{}); err != nil {
		panic("Failed to migrate Address model")
	}
	// Return the database connection
	return db
}

func main() {
	// Initialize the database
	db := initDB()
	// Initialize the router
	router := gin.Default()
	// Initialize the customer controller
	customerRepository := repository.NewCustomerRepository(db)
	customerService := service.NewCustomerService(customerRepository)
	controller.NewCustomerController(router, customerService)
	// Initialize the address controller
	addressRepository := repository.NewAddressRepository(db)
	addressService := service.NewAddressService(addressRepository)
	controller.NewAddressController(router, addressService)
	// Run the server
	router.Run()
}
