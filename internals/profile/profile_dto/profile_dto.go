package profiledto

type ProfileDTO struct {
	ID      uint       `json:"id"`
	Name    string     `json:"name"`
	Email   string     `json:"email"`
	Role    string     `json:"role"`
	Address AddressDTO `json:"address"`
}

type AddressDTO struct {
	ID          uint   `json:"id"`
	AddressLine string `json:"street" binding:"required"`
	City        string `json:"city" binding:"required"`
	State       string `json:"state" binding:"required"`
	Phone       string `json:"phone" binding:"required"`
	PostalCode  string `json:"zip_code" binding:"required"`
	Country     string `json:"country" binding:"required"`
}
