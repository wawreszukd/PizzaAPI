package handlers_test

import (
	"PizzeriaAPI/handlers"
	"github.com/DATA-DOG/go-sqlmock"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetAllPizzasHandler(t *testing.T) {
	db, _, _ := sqlmock.New()
	defer db.Close()

	handler := handlers.Handler{DB: db}

	router := gin.Default()
	router.GET("/pizzas", handler.GetAllPizzas)

	req, _ := http.NewRequest("GET", "/pizzas", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestGetPizzaByIdHandler(t *testing.T) {
	db, _, _ := sqlmock.New()
	defer db.Close()

	handler := handlers.Handler{DB: db}

	router := gin.Default()
	router.GET("/pizzas/:id", handler.GetPizzaById)

	req, _ := http.NewRequest("GET", "/pizzas/1", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestDeletePizzaHandler(t *testing.T) {
	db, _, _ := sqlmock.New()
	defer db.Close()

	handler := handlers.Handler{DB: db}

	router := gin.Default()
	router.DELETE("/pizzas/:id", handler.DeletePizza)

	req, _ := http.NewRequest("DELETE", "/pizzas/1", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestUpdatePizzaHandler(t *testing.T) {
	db, _, _ := sqlmock.New()
	defer db.Close()

	handler := handlers.Handler{DB: db}

	router := gin.Default()
	router.PUT("/pizzas/:id", handler.UpdatePizza)

	req, _ := http.NewRequest("PUT", "/pizzas/1", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestInsertPizzaHandler(t *testing.T) {
	db, _, _ := sqlmock.New()
	defer db.Close()

	handler := handlers.Handler{DB: db}

	router := gin.Default()
	router.POST("/pizzas", handler.InsertPizza)

	req, _ := http.NewRequest("POST", "/pizzas", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}
