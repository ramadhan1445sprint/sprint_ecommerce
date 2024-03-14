package svc

import (
	"testing"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/ramadhan1445sprint/sprint_ecommerce/config"
	"github.com/ramadhan1445sprint/sprint_ecommerce/entity"
	"github.com/ramadhan1445sprint/sprint_ecommerce/pkg/database"
	"github.com/ramadhan1445sprint/sprint_ecommerce/repo"
)

func strPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}

func TestCreateBankAccount(t *testing.T) {
	config.LoadConfig("../.env")

	db, err := database.NewDatabase()
	if err != nil {
		t.Fatalf("failed to create a database connection: %v", err)
	}

	bankAccountRepo := repo.NewBankAccountRepo(db)
	bankAccountSvc := NewBankAccounthSvc(bankAccountRepo)

	testCases := []struct {
		name        		string
		input       		entity.BankAccountCreateRequest
		userID				int
		statusExpected 		int
		errExpected 		bool
	}{
		{"Test create bank success", entity.BankAccountCreateRequest{Name: strPtr("svc test"), AccountName: strPtr("svc taest name"), AccountNumber: intPtr(23123)}, 1, 200, false},
		{"Test create bank failed", entity.BankAccountCreateRequest{Name: nil, AccountName: strPtr("svc taest name"), AccountNumber: intPtr(23123)}, 1, 400, true},
		{"Test create bank failed", entity.BankAccountCreateRequest{Name: strPtr("svc test"), AccountName: nil, AccountNumber: intPtr(23123)}, 1, 400, true},
		{"Test create bank failed", entity.BankAccountCreateRequest{Name: strPtr("svc test"), AccountName: strPtr("svc taest name"), AccountNumber: nil}, 1, 400, true},
		{"Test create bank failed", entity.BankAccountCreateRequest{Name: strPtr(""), AccountName: strPtr("svc taest name"), AccountNumber: intPtr(23123)}, 1, 400, true},
		{"Test create bank failed", entity.BankAccountCreateRequest{Name: strPtr("svc test"), AccountName: strPtr(""), AccountNumber: intPtr(23123)}, 1, 400, true},
		{"Test create bank failed", entity.BankAccountCreateRequest{Name: strPtr("svc test"), AccountName: strPtr("dadawd"), AccountNumber: intPtr(0)}, 1, 400, true},
		{"Test create bank failed", entity.BankAccountCreateRequest{Name: strPtr("svc test iadiajwdijaidjaiddawjiawjd"), AccountName: strPtr("dadawd"), AccountNumber: intPtr(0)}, 1, 400, true},
		{"Test create bank failed", entity.BankAccountCreateRequest{Name: strPtr("svc test"), AccountName: strPtr("dkoakdoak oakwdoawkdo koakwdoakwdkkdoaw"), AccountNumber: intPtr(0)}, 1, 400, true},
		{"Test create bank failed", entity.BankAccountCreateRequest{Name: strPtr("svc test"), AccountName: strPtr("dadawd"), AccountNumber: intPtr(123)}, 1, 400, true},
		{"Test create bank failed", entity.BankAccountCreateRequest{Name: strPtr("svc test"), AccountName: strPtr("dadawd"), AccountNumber: intPtr(1234567890123456)}, 1, 400, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			status, err := bankAccountSvc.CreateBankAccount(&tc.input, tc.userID)

			if tc.errExpected {
				if err == nil {
					t.Errorf("Expected error, but no error")
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, but got: %v", err)
				}

				if status != 200 {
					t.Errorf("expected %d, but got %d", tc.statusExpected, status)
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

	bankAccountRepo := repo.NewBankAccountRepo(db)
	bankAccountSvc := NewBankAccounthSvc(bankAccountRepo)

	testCases := []struct {
		name        string
		input       int
		errExpected bool
	}{
		{"Test get bank account success", 1, false},
		{"Test get bank account fail", 99999, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, status, err := bankAccountSvc.GetBankAccount(tc.input)

			if tc.errExpected {
				if err == nil {
					t.Errorf("Expected error, but no error")
				}
			} else {

				if err != nil {
					t.Errorf("Expected no error, but got: %v", err)
				}

				if status != 200 {
					t.Errorf("expected %d, but got %d", 200, status)
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

	bankAccountRepo := repo.NewBankAccountRepo(db)
	bankAccountSvc := NewBankAccounthSvc(bankAccountRepo)

	testCases := []struct {
		name        string
		input       entity.BankAccountUpdateRequest
		errExpected bool
	}{
		{"Test update bank account success", entity.BankAccountUpdateRequest{ID: intPtr(5), Name: strPtr("update svc 2"), AccountName: strPtr("dadang svc 2"), AccountNumber: intPtr(123454)}, false},
		{"Test failed bank failed", entity.BankAccountUpdateRequest{ID: intPtr(5), Name: nil, AccountName: strPtr("svc taest name"), AccountNumber: intPtr(23123)}, true},
		{"Test failed bank failed", entity.BankAccountUpdateRequest{ID: intPtr(5), Name: strPtr("svc test"), AccountName: nil, AccountNumber: intPtr(23123)}, true},
		{"Test failed bank failed", entity.BankAccountUpdateRequest{ID: intPtr(5), Name: strPtr("svc test"), AccountName: strPtr("svc taest name"), AccountNumber: nil}, true},
		{"Test failed bank failed", entity.BankAccountUpdateRequest{ID: intPtr(5), Name: strPtr(""), AccountName: strPtr("svc taest name"), AccountNumber: intPtr(23123)}, true},
		{"Test failed bank failed", entity.BankAccountUpdateRequest{ID: intPtr(5), Name: strPtr("svc test"), AccountName: strPtr(""), AccountNumber: intPtr(23123)}, true},
		{"Test update bank failed", entity.BankAccountUpdateRequest{ID: intPtr(5), Name: strPtr("svc test"), AccountName: strPtr("dadawd"), AccountNumber: intPtr(0)}, true},
		{"Test update bank failed", entity.BankAccountUpdateRequest{ID: intPtr(5), Name: strPtr("svc test iadiajwdijaidjaiddawjiawjd"), AccountName: strPtr("dadawd"), AccountNumber: intPtr(0)}, true},
		{"Test update bank failed", entity.BankAccountUpdateRequest{ID: intPtr(5), Name: strPtr("svc test"), AccountName: strPtr("dkoakdoak oakwdoawkdo koakwdoakwdkkdoaw"), AccountNumber: intPtr(0)}, true},
		{"Test update bank failed", entity.BankAccountUpdateRequest{ID: intPtr(5), Name: strPtr("svc test"), AccountName: strPtr("dadawd"), AccountNumber: intPtr(123)}, true},
		{"Test update bank failed", entity.BankAccountUpdateRequest{ID: intPtr(5), Name: strPtr("svc test"), AccountName: strPtr("dadawd"), AccountNumber: intPtr(1234567890123456)}, true},
		{"Test update bank failed", entity.BankAccountUpdateRequest{ID: intPtr(4234234), Name: strPtr("svc test"), AccountName: strPtr("dadawd"), AccountNumber: intPtr(1234567890123456)}, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			status, err := bankAccountSvc.UpdateBankAccount(&tc.input)

			if tc.errExpected {
				if err == nil {
					t.Errorf("Expected error, but no error")
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, but got: %v", err)
				}

				if status != 200 {
					t.Errorf("expected %d, but got %d", 200, status)
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

	bankAccountRepo := repo.NewBankAccountRepo(db)
	bankAccountSvc := NewBankAccounthSvc(bankAccountRepo)

	testCases := []struct {
		name        string
		input       int
		errExpected bool
	}{
		{"Test delete bank account success", 9, false},
		{"Test delete bank account failed", 99999, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			status, err := bankAccountSvc.DeleteBankAccount(tc.input)

			if tc.errExpected {
				if err == nil {
					t.Errorf("Expected error, but no error")
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, but got: %v", err)
				}

				if status != 200 {
					t.Errorf("expected %d, but got %d", 200, status)
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

	paymentRepo := repo.NewPaymentRepo(db)
	paymentSvc := NewPaymentSvc(paymentRepo)

	testCases := []struct {
		name        string
		input       entity.PaymentCreateRequest
		errExpected bool
	}{
		{"Test create payment success", entity.PaymentCreateRequest{ProductID: intPtr(1), BankAccountID: intPtr(6), PaymentProofImgUrl: strPtr("https://example.com/image.jpg"), Quantity: intPtr(10)}, false},
		{"Test create payment failed 1", entity.PaymentCreateRequest{ProductID: intPtr(1), BankAccountID: nil, PaymentProofImgUrl: strPtr("https://example.com/image.jpg"), Quantity: intPtr(10)}, true},
		{"Test create payment failed 2", entity.PaymentCreateRequest{ProductID: nil, BankAccountID: intPtr(6), PaymentProofImgUrl: strPtr("https://example.com/image.jpg"), Quantity: intPtr(10)}, true},
		{"Test create payment failed 3", entity.PaymentCreateRequest{ProductID: intPtr(1), BankAccountID: intPtr(6), PaymentProofImgUrl: nil, Quantity: intPtr(10)}, true},
		{"Test create payment failed 4", entity.PaymentCreateRequest{ProductID: intPtr(1), BankAccountID: intPtr(6), PaymentProofImgUrl: strPtr("https://example.com/image.jpg"), Quantity: nil}, true},
		{"Test create payment failed 5", entity.PaymentCreateRequest{ProductID: intPtr(1), BankAccountID: intPtr(6), PaymentProofImgUrl: strPtr("https://example.com/image.jpg"), Quantity: intPtr(0)}, true},
		{"Test create payment failed 6", entity.PaymentCreateRequest{ProductID: intPtr(1), BankAccountID: intPtr(6), PaymentProofImgUrl: strPtr("owkeokweo"), Quantity: intPtr(0)}, true},
		{"Test create payment failed 7", entity.PaymentCreateRequest{ProductID: intPtr(9999), BankAccountID: intPtr(6), PaymentProofImgUrl: strPtr("https://example.com/image.jpg"), Quantity: intPtr(10)}, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			status, err := paymentSvc.CreatePayment(&tc.input)

			if tc.errExpected {
				if err == nil {
					t.Errorf("Expected error but no error")
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error, but got: %v", err)
				}

				if status != 200 {
					t.Errorf("expected %d, but got %d", 200, status)
				}
			}
		})
	}
}
