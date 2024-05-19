package handlers

import (
	"PizzeriaAPI/models"
	"PizzeriaAPI/storage"
	"database/sql"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Handler struct is used to handle database operations.
type Handler struct {
	DB *sql.DB
}

// GetAllPizzas is a handler function that retrieves all pizzas from the database.
// It writes the result as JSON to the response body.
func (h *Handler) GetAllPizzas(c *gin.Context) {
	pizzas := storage.GetAllPizzas(h.DB)
	c.IndentedJSON(200, pizzas)
}

// GetPizzaById is a handler function that retrieves a pizza by its ID from the database.
// It writes the result as JSON to the response body.
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

// DeletePizza is a handler function that deletes a pizza by its ID from the database.
// It writes the result as JSON to the response body.
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

// UpdatePizza is a handler function that updates a pizza by its ID in the database.
// It writes the result as JSON to the response body.
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

// InsertPizza is a handler function that inserts a new pizza into the database.
// It writes the result as JSON to the response body.
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