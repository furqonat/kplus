package dto

type ResponseSignInDto struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type SignInDto struct {
	PhoneOrEmail string `json:"username"`
	Password     string `json:"password"`
}

type QuerySignIn struct {
	Provider string `json:"provider"`
}

type SignUpDto struct {
	Role     *string `json:"role,omitempty"`
	Phone    string  `json:"phone"`
	Email    *string `json:"email,omitempty"`
	Password string  `json:"password"`
}
