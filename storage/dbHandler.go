package storage

import (
	"PizzeriaAPI/models"
	"database/sql"
	"log"
	"os"
	"github.com/joho/godotenv"
)

// Connect establishes a connection to the database using environment variables for configuration.
// It loads the .env file, retrieves the necessary variables, and opens a connection to the database.
// It also checks the connection and creates the pizzas table if it doesn't exist.
// It returns a pointer to the sql.DB object representing the database connection.
func Connect() *sql.DB {
	// Load environment variables from .env file
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatalf("Error loading .env file: %s", err)
	}
	// Retrieve environment variables
	user:= os.Getenv("USER")
	password:= os.Getenv("PASSWORD")
	host:= os.Getenv("HOST")
	port:= os.Getenv("PORT")
	dbname:= os.Getenv("DBNAME")
	// Construct connection string
	connStr := "user=" + user + " password=" + password + " host=" + host + " port=" + port + " dbname=" + dbname + " sslmode=disable"
	// Open connection to database
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	// Create pizzas table if it doesn't exist
	createTable(db)
	return db
}

// createTable creates the pizzas table in the database if it doesn't exist.
// It executes a SQL query to create the table and logs any error that occurs.
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

// InsertPizza inserts a new pizza into the database.
// It executes a SQL query to insert the pizza and returns the ID of the new pizza.
// It logs any error that occurs.
func InsertPizza(db *sql.DB, pizza models.Pizza) int {
	query := `INSERT INTO pizzas (name, price, dough) VALUES ($1, $2, $3) RETURNING id`
	var id int
	err := db.QueryRow(query, pizza.Name, pizza.Price, pizza.Dough).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}
	return id
}

// GetAllPizzas retrieves all pizzas from the database.
// It executes a SQL query to select all pizzas and returns a slice of Pizza objects.
// It logs any error that occurs.
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

// GetPizzaById retrieves a pizza by its ID from the database.
// It executes a SQL query to select the pizza and returns a Pizza object.
// It logs any error that occurs.
func GetPizzaById(db *sql.DB, id int) models.Pizza {
	query := `SELECT * FROM pizzas WHERE id = $1`
	var pizza models.Pizza
	err := db.QueryRow(query, id).Scan(&pizza.ID, &pizza.Name, &pizza.Price, &pizza.Dough)
	if err != nil {
		log.Fatal(err)
	}
	return pizza
}

// DeletePizza deletes a pizza by its ID from the database.
// It executes a SQL query to delete the pizza and returns a Pizza object representing the deleted pizza.
// It logs any error that occurs.
func DeletePizza(db *sql.DB, id int) models.Pizza {
	query := `DELETE FROM pizzas WHERE id = $1 RETURNING *`
	var pizza models.Pizza
	err := db.QueryRow(query, id).Scan(&pizza.ID, &pizza.Name, &pizza.Price, &pizza.Dough)
	if err != nil {
		log.Fatal(err)
	}
	return pizza
}

// UpdatePizza updates a pizza by its ID in the database.
// It executes a SQL query to update the pizza and returns a Pizza object representing the updated pizza.
// It logs any error that occurs.
func UpdatePizza(db *sql.DB, id int, pizza models.Pizza) models.Pizza {
	query := `UPDATE pizzas SET name = $1, price = $2, dough = $3 WHERE id = $4 RETURNING *`
	var updatedPizza models.Pizza
	err := db.QueryRow(query, pizza.Name, pizza.Price, pizza.Dough, id).Scan(&updatedPizza.ID, &updatedPizza.Name, &updatedPizza.Price, &updatedPizza.Dough)
	if err != nil {
		log.Fatal(err)
	}
	return updatedPizza
}