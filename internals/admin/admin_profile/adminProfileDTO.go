package adminprofile

type ProfileDTO struct {
	ID      uint       `json:"id"`
	Name    string     `json:"name"`
	Email   string     `json:"email"`
	Role    string     `json:"role"`
	Address AddressDTO `json:"address"`
}

type AddressDTO struct {
	ID          uint   `json:"id"`
	AddressLine string `json:"street"`
	City        string `json:"city"`
	State       string `json:"state"`
	Phone       string `json:"phone"`
	PostalCode  string `json:"zip_code"`
	Country     string `json:"country"`
}
