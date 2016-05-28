package models

//Supplier Structure
type Supplier struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Link        string    `json:"link"`
	Info        []Contact `json:"info"`
}

//Contact Structure
type Contact struct {
	Service string `json:"service"`
	Address string `json:"address"`
}
