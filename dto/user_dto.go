package dto

type UserDto struct {
	ID       int            `json:"id"`
	Phone    string         `json:"phone,omitempty"`
	Email    string         `json:"email,omitempty"`
	Password string         `json:"password,omitempty"`
	Role     string         `json:"role,omitempty"`
	Details  UserDetailsDto `json:"details,omitempty"`
}

type UserDetailsDto struct {
	ID                   int     `json:"id"`
	Nik                  *string `json:"nik"`
	UserID               *string `json:"user_id"`
	FullName             *string `json:"full_name"`
	LegalName            *string `json:"legal_name"`
	PlaceOfBirth         *string `json:"place_of_birth"`
	DateOfBirth          *string `json:"date_of_birth"`
	Salary               *string `json:"salary"`
	Selfie               *string `json:"selfie"`
	SelfieWithNationalID *string `json:"selfie_with_national_id"`
	NationalIdImage      *string `json:"national_id_image"`
}
