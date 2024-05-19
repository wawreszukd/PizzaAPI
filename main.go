package main

import (
	"PizzeriaAPI/handlers"
	"PizzeriaAPI/storage"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	db := storage.Connect()
	router := gin.Default()
	handler := handlers.Handler{DB: db}
	router.GET("/pizza", handler.GetAllPizzas)
	router.GET("/pizza/:id", handler.GetPizzaById)
	router.POST("/pizza", handler.InsertPizza)
	router.DELETE("/pizza/:id", handler.DeletePizza)
	router.PUT("/pizza/:id", handler.UpdatePizza)
	defer db.Close()
	err := router.Run(":8080")
	if err != nil {
		return
	}

}
