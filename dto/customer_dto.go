package dto

type CustomerDto struct {
	ID       string `json:"id"`
	Nik      string `json:"nik"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
