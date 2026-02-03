package dtos

type SearchAddressRequest struct {
	Name    string
	City    string
	Country string
	Limit   int
}

type SearchAddressResponse struct {
	DisplayName string `json:"display_name"`
}
