package entity

type Payment struct {
	PaymentProofImgUrl string `json:"paymentProofImageUrl"`
	ID                 int   `json:"id"`
	ProductID          int   `json:"productId"`
	BankAccountID      int   `json:"bankAccountId"`
	Quantity           int   `json:"quantity"`
}

type PaymentCreateRequest struct {
	PaymentProofImgUrl *string `json:"paymentProofImageUrl"`
	ProductID          *int   `json:"productId"`
	BankAccountID      *int   `json:"bankAccountId"`
	Quantity           *int   `json:"quantity"`
}
