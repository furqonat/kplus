package dto

type TransactionDto struct {
	ID             string `json:"id"`
	ContractNumber string `json:"contract_number"`
	UserID         int    `json:"user_id"`
	OTR            string `json:"otr"`
	Fee            string `json:"fee"`
	Installment    string `json:"installment"`
	Interest       string `json:"interest"`
	Status         string `json:"status"`
	AssetName      string `json:"asset_name"`
}
