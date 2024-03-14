package entity

type BankAccount struct {
	Name          string `json:"bankName" db:"bank_name"` 
	AccountName   string `json:"bankAccountName" db:"account_name"`
	UserID        int `db:"user_id"`
	ID            int `json:"bankAccountId" db:"id"`
	AccountNumber int `json:"bankAccountNumber" db:"account_number"`
}

type BankAccountCreateRequest struct {
	Name          *string `json:"bankName"`
	AccountName   *string `json:"bankAccountName"`
	AccountNumber *int   `json:"bankAccountNumber"`
}

type BankAccountCreateResponse struct {
	Message string `json:"message"`
}

type BankAccountGetResponse struct {
	Name          string `json:"bankName"`
	AccountName   string `json:"bankAccountName"`
	ID            int   `json:"bankAccountId"`
	AccountNumber int   `json:"bankAccountNumber"`
}

type BankAccountUpdateRequest struct {
	Name          *string `json:"bankName"`
	AccountName   *string `json:"bankAccountName"`
	ID            *int   `json:"bankAccountId"`
	AccountNumber *int   `json:"bankAccountNumber"`
}
