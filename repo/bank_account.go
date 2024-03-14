package repo

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/ramadhan1445sprint/sprint_ecommerce/entity"
)

type BankAccountRepoInterface interface {
	GetBankAccount(userId int) ([]entity.BankAccount, error)
	CreateBankAccount(bankAccount *entity.BankAccount) error
	UpdateBankAccount(bankAccount *entity.BankAccount) error
	DeleteBankAccount(bankAccountID int) error
}

func NewBankAccountRepo(db *sqlx.DB) BankAccountRepoInterface {
	return &bankAccountRepo{db: db}
}

type bankAccountRepo struct {
	db *sqlx.DB
}

func (r *bankAccountRepo) GetBankAccount(userId int) ([]entity.BankAccount, error) {
	var res []entity.BankAccount

	r.db.Select(&res, "SELECT id, bank_name, account_name, account_number from bank_account where user_id = $1", userId)

	if len(res) == 0 {
		return nil, errors.New("bank account not found")
	}
	return res, nil
}

func (r *bankAccountRepo) CreateBankAccount(bankAccount *entity.BankAccount) error {
	_, err := r.db.Exec("insert into bank_account (user_id, bank_name, account_name, account_number) values ($1, $2, $3, $4)", bankAccount.UserID, bankAccount.Name, bankAccount.AccountName, bankAccount.AccountNumber)

	if err != nil {
		return err
	}

	return nil
}

func (r *bankAccountRepo) UpdateBankAccount(bankAccount *entity.BankAccount) error {
	res, _ := r.db.Exec("UPDATE bank_account SET bank_name = $1, account_name = $2, account_number = $3 where id = $4", bankAccount.Name, bankAccount.AccountName, bankAccount.AccountNumber, bankAccount.ID)

	rowsEffected, _ := res.RowsAffected()

	if rowsEffected == 0 {
		return errors.New("bank account not found")
	}

	return nil
}

func (r *bankAccountRepo) DeleteBankAccount(bankAccountID int) error {
	res, _ := r.db.Exec("delete from bank_account where id = $1", bankAccountID)

	rowsEffected, _ := res.RowsAffected()

	if rowsEffected == 0 {
		return errors.New("bank account not found")
	}

	return nil
}
