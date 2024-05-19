# Pizza API
This is a simple API for pizza management. It allows to create, read, update and delete pizzas.

## Running the application
To run the application, execute the following command:
```bash
go run main.go
```
## Application endpoints
The application has the following endpoints:
### Get all
``` http request 
### GET all pizzas
GET localhost:8080/pizza
###
```
### Get one
```http request
### GET pizza id 1
GET localhost:8080/pizza/3
```
### Post
```http request 
### POST new pizza
POST localhost:8080/pizza
content-type: application/json

{
"name": "Pizza 1",
"dough": "white dough",
"price": 10.0
}
```
### Delete
```http request
### Delete pizza id 2
DELETE localhost:8080/pizza/2
```
### Put
```http request
### PUT pizza id 2
PUT localhost:8080/pizza/3
content-type: application/json

{
"name": "Pizza 333",
"dough": "white dough",
"price": 20.0
}
```