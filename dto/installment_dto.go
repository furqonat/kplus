package dto

type InstallmentDto struct {
	ID             int            `json:"id"`
	TransactionID  int            `json:"transaction_id"`
	Installment    string         `json:"installment"`
	DueDate        string         `json:"due_date"`
	PaidDate       string         `json:"paid_date"`
	Period         int            `json:"period"`
	Status         string         `json:"status"`
	CreatedAt      string         `json:"created_at"`
	TransactionDto TransactionDto `json:"transaction"`
}
