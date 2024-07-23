package dto

type LoanDto struct {
	ID     string `json:"id"`
	UserID int    `json:"user_id"`
	Limit  int    `json:"limit"`
	Tenor  int    `json:"tenor"`
}
