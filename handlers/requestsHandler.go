package handlers

import (
	"PizzeriaAPI/models"
	"PizzeriaAPI/storage"
	"database/sql"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Handler struct {
	DB *sql.DB
}

func (h *Handler) GetAllPizzas(c *gin.Context) {
	pizzas := storage.GetAllPizzas(h.DB)
	c.IndentedJSON(200, pizzas)
}
func (h *Handler) GetPizzaById(c *gin.Context) {
	id := c.Param("id")
	i, err := strconv.Atoi(id)

	if err != nil {

		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	pizza := storage.GetPizzaById(h.DB, i)
	c.IndentedJSON(200, pizza)
}
func (h *Handler) DeletePizza(c *gin.Context) {
	id := c.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	pizza := storage.DeletePizza(h.DB, i)
	c.IndentedJSON(200, pizza)
}
func (h *Handler) UpdatePizza(c *gin.Context) {
	id := c.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	var pizzaData models.Pizza
	if err := c.BindJSON(&pizzaData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	pizza := storage.UpdatePizza(h.DB, i, pizzaData)
	c.IndentedJSON(200, pizza)

}
func (h *Handler) InsertPizza(c *gin.Context) {
	var pizza models.Pizza
	err := c.BindJSON(&pizza)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}
	id := storage.InsertPizza(h.DB, pizza)
	pizza.ID = id
	c.IndentedJSON(200, pizza)
}
