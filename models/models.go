package models

// scema of the rice table
type Rice struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Location string    `json:"location"`
}