package storage

import (
	"PizzeriaAPI/models"
	"database/sql"
	"log"
)

func Connect() *sql.DB {
	connStr := "postgres://postgres:secret@localhost:5432/gopostgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	createTable(db)
	return db
}

func createTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS pizzas (
			id SERIAL PRIMARY KEY,
			name TEXT,
			price FLOAT,
			dough TEXT
		);
	`
	_, err := db.Exec(query)
	if err != nil {
		return
	}
}

func InsertPizza(db *sql.DB, pizza models.Pizza) int {
	query := `INSERT INTO pizzas (name, price, dough) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := db.QueryRow(query, pizza.Name, pizza.Price, pizza.Dough).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	return id
}

func GetAllPizzas(db *sql.DB) []models.Pizza {
	query := `SELECT * FROM pizzas`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var pizzas []models.Pizza
	for rows.Next() {
		var pizza models.Pizza
		err := rows.Scan(&pizza.ID, &pizza.Name, &pizza.Price, &pizza.Dough)
		if err != nil {
			log.Fatal(err)
		}
		pizzas = append(pizzas, pizza)
	}
	return pizzas
}
func GetPizzaById(db *sql.DB, id int) models.Pizza {
	query := `SELECT * FROM pizzas WHERE id = $1`
	var pizza models.Pizza
	err := db.QueryRow(query, id).Scan(&pizza.ID, &pizza.Name, &pizza.Price, &pizza.Dough)
	if err != nil {
		log.Fatal(err)
	}
	return pizza
}
func DeletePizza(db *sql.DB, id int) models.Pizza {
	query := `DELETE FROM pizzas WHERE id = $1 RETURNING *`
	var pizza models.Pizza
	err := db.QueryRow(query, id).Scan(&pizza.ID, &pizza.Name, &pizza.Price, &pizza.Dough)
	if err != nil {
		log.Fatal(err)
	}
	return pizza
}
func UpdatePizza(db *sql.DB, id int, pizza models.Pizza) models.Pizza {
	query := `UPDATE pizzas SET name = $1, price = $2, dough = $3 WHERE id = $4 RETURNING *`
	var updatedPizza models.Pizza
	err := db.QueryRow(query, pizza.Name, pizza.Price, pizza.Dough, id).Scan(&updatedPizza.ID, &updatedPizza.Name, &updatedPizza.Price, &updatedPizza.Dough)
	if err != nil {
		log.Fatal(err)
	}
	return updatedPizza
}
