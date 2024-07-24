package dto

type LoanDto struct {
	ID     int     `json:"id,omitempty"`
	UserID int     `json:"user_id,omitempty"`
	Limit  float64 `json:"limit"`
	Tenor  int     `json:"tenor,omitempty"`
}

type CreateLoanDto struct {
	Limit float64 `json:"limit"`
	Tenor int     `json:"tenor"`
}
