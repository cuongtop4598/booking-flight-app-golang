package requests

type Customer struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	LicenseId   string `json:"license_id"`
	PhoneNumber string `json:"phone" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Active      bool   `json:"active"`
}

type ListCustomerRequests struct {
	Active bool `json:"active"`
}
