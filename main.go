package main

import (
	"PizzeriaAPI/handlers"
	"PizzeriaAPI/storage"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// Connect to the database
	db := storage.Connect()

	router := gin.Default()
	// Create a new handler to work with database
	handler := handlers.Handler{DB: db}
	// Routers
	router.GET("/pizza", handler.GetAllPizzas)
	router.GET("/pizza/:id", handler.GetPizzaById)
	router.POST("/pizza", handler.InsertPizza)
	router.DELETE("/pizza/:id", handler.DeletePizza)
	router.PUT("/pizza/:id", handler.UpdatePizza)
	// Close the database connection
	defer db.Close()
	err := router.Run(":8080")
	if err != nil {
		return
	}

}
