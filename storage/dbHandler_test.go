package storage_test

import (
	"PizzeriaAPI/models"
	"PizzeriaAPI/storage"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestInsertPizza(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	mock.ExpectQuery("INSERT INTO pizzas").WithArgs("Margherita", 10.5, "Thin").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	pizza := models.Pizza{Name: "Margherita", Price: 10.5, Dough: "Thin"}
	id := storage.InsertPizza(db, pizza)

	assert.Equal(t, 1, id)
}

func TestGetAllPizzas(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "price", "dough"}).
		AddRow(1, "Margherita", 10.5, "Thin").
		AddRow(2, "Pepperoni", 12.5, "Thick")

	mock.ExpectQuery("SELECT \\* FROM pizzas").WillReturnRows(rows)

	pizzas := storage.GetAllPizzas(db)

	assert.Equal(t, 2, len(pizzas))
	assert.Equal(t, "Margherita", pizzas[0].Name)
	assert.Equal(t, "Pepperoni", pizzas[1].Name)
}

func TestGetPizzaById(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "price", "dough"}).
		AddRow(1, "Margherita", 10.5, "Thin")

	mock.ExpectQuery("SELECT \\* FROM pizzas WHERE id = \\$1").WithArgs(1).WillReturnRows(rows)

	pizza := storage.GetPizzaById(db, 1)

	assert.Equal(t, "Margherita", pizza.Name)
}

func TestDeletePizza(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "price", "dough"}).
		AddRow(1, "Margherita", 10.5, "Thin")

	mock.ExpectQuery("DELETE FROM pizzas WHERE id = \\$1 RETURNING \\*").WithArgs(1).WillReturnRows(rows)

	pizza := storage.DeletePizza(db, 1)

	assert.Equal(t, "Margherita", pizza.Name)
}

func TestUpdatePizza(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "price", "dough"}).
		AddRow(1, "Margherita", 11.5, "Thin")

	mock.ExpectQuery("UPDATE pizzas SET name = \\$1, price = \\$2, dough = \\$3 WHERE id = \\$4 RETURNING \\*").
		WithArgs("Margherita", 11.5, "Thin", 1).WillReturnRows(rows)

	pizza := models.Pizza{Name: "Margherita", Price: 11.5, Dough: "Thin"}
	updatedPizza := storage.UpdatePizza(db, 1, pizza)

	assert.Equal(t, float32(11.5), updatedPizza.Price)
}
