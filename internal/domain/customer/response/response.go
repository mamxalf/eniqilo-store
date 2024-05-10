package response

type CustomerCreatedResponse struct {
	UserID      string `json:"userId"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
}
