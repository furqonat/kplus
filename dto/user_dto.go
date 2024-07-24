package dto

type UserDto struct {
	ID       int            `json:"id"`
	Phone    string         `json:"phone,omitempty"`
	Email    *string        `json:"email,omitempty"`
	Password string         `json:"password,omitempty"`
	Role     string         `json:"role,omitempty"`
	Details  UserDetailsDto `json:"details,omitempty"`
}

type UserDetailsDto struct {
	ID                   *int    `json:"id,omitempty"`
	Nik                  *string `json:"nik,omitempty"`
	UserID               *int    `json:"user_id,omitempty"`
	FullName             *string `json:"full_name,omitempty"`
	LegalName            *string `json:"legal_name,omitempty"`
	PlaceOfBirth         *string `json:"place_of_birth,omitempty"`
	DateOfBirth          *string `json:"date_of_birth,omitempty"`
	Salary               *string `json:"salary,omitempty"`
	Selfie               *string `json:"selfie,omitempty"`
	SelfieWithNationalID *string `json:"selfie_with_national_id,omitempty"`
	NationalIdImage      *string `json:"national_id_image,omitempty"`
}
