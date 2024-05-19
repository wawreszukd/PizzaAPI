package models

type Pizza struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
	Dough string  `json:"dough"`
}
