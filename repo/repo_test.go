package repo

import (
	"testing"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/ramadhan1445sprint/sprint_ecommerce/config"
	"github.com/ramadhan1445sprint/sprint_ecommerce/entity"
	"github.com/ramadhan1445sprint/sprint_ecommerce/pkg/database"
)

func TestCreateBankAccount(t *testing.T) {
	config.LoadConfig("../.env")

	db, err := database.NewDatabase()
	if err != nil {
		t.Fatalf("failed to create a database connection: %v", err)
	}

	bankAccountRepo := NewBankAccountRepo(db)

	testCases := []struct {
		name        string
		input       entity.BankAccount
		errExpected bool
	}{
		{"Test create bank", entity.BankAccount{UserID: 1, Name: "BCA", AccountName: "Ilham Nuryanto", AccountNumber: 1234567}, false},
		{"Test create bank 2", entity.BankAccount{UserID: 1, Name: "Mandiri", AccountName: "Ilham Nuryanto", AccountNumber: 432145}, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := bankAccountRepo.CreateBankAccount(&tc.input)

			if tc.errExpected {
				if err == nil {
					t.Errorf("Expected error, but no error")
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, but got: %v", err)
				}
			}
		})
	}

}
func TestGetBankAccount(t *testing.T) {
	config.LoadConfig("../.env")

	db, err := database.NewDatabase()
	if err != nil {
		t.Fatalf("failed to create a database connection: %v", err)
	}

	bankAccountRepo := NewBankAccountRepo(db)

	testCases := []struct {
		name        string
		input       int
		errExpected bool
	}{
		{"Test get bank account found", 1, false},
		{"Test get bank account not found", 99999, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := bankAccountRepo.GetBankAccount(tc.input)

			if tc.errExpected {
				if err == nil {
					t.Errorf("Expected error, but no error")
				}
			} else {
				if len(res) == 0 {
					t.Errorf("Expected exist, but got: %d", len(res))
				}

				if err != nil {
					t.Errorf("Expected no error, but got: %v", err)
				}
			}
		})
	}
}

func TestUpdateBankAccount(t *testing.T) {
	config.LoadConfig("../.env")

	db, err := database.NewDatabase()
	if err != nil {
		t.Fatalf("failed to create a database connection: %v", err)
	}

	bankAccountRepo := NewBankAccountRepo(db)

	testCases := []struct {
		name        string
		input       entity.BankAccount
		errExpected bool
	}{
		{"Test update bank account success", entity.BankAccount{ID: 6, Name: "BCA updated", AccountName: "ilham updated", AccountNumber: 123123}, false},
		{"Test update bank account failed", entity.BankAccount{ID: 99999, Name: "BCA", AccountName: "dadang", AccountNumber: 1234}, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := bankAccountRepo.UpdateBankAccount(&tc.input)

			if tc.errExpected {
				if err == nil {
					t.Errorf("Expected error, but no error")
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, but got: %v", err)
				}
			}
		})
	}
}
func TestDeleteBankAccount(t *testing.T) {
	config.LoadConfig("../.env")

	db, err := database.NewDatabase()
	if err != nil {
		t.Fatalf("failed to create a database connection: %v", err)
	}

	bankAccountRepo := NewBankAccountRepo(db)

	testCases := []struct {
		name        string
		input       int
		errExpected bool
	}{
		{"Test delete bank account success", 4, false},
		{"Test delete bank account failed", 9999, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := bankAccountRepo.DeleteBankAccount(tc.input)

			if tc.errExpected {
				if err == nil {
					t.Errorf("Expected error, but no error")
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, but got: %v", err)
				}
			}
		})
	}
}
func TestCreatePayment(t *testing.T) {
	config.LoadConfig("../.env")

	db, err := database.NewDatabase()
	if err != nil {
		t.Fatalf("failed to create a database connection: %v", err)
	}

	paymentRepo := NewPaymentRepo(db)

	testCases := []struct {
		name        string
		input       entity.Payment
		errExpected bool
	}{
		{"Test create payment success", entity.Payment{ProductID: 1, BankAccountID: 5, PaymentProofImgUrl: "url", Quantity: 10}, false},
		{"Test create payment failed", entity.Payment{ProductID: 1, BankAccountID: 9999, PaymentProofImgUrl: "url", Quantity: 10}, true},
		{"Test create payment failed", entity.Payment{ProductID: 9999, BankAccountID: 4, PaymentProofImgUrl: "url", Quantity: 10}, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := paymentRepo.CreatePayment(&tc.input)

			if tc.errExpected {
				if err == nil {
					t.Errorf("Expected error but no error")
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, but got: %v", err)
				}
			}
		})
	}
}
