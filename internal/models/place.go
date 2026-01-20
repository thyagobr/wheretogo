package models

type Place struct {
	ID			int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Country string `json:"country"`
	City    string `json:"city"`
	//Description string `json:"description"`
}
