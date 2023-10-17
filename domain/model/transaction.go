package model

import (
	"time"
	uuid "github.com/satori/go.uuid"
	"github.com/asaskevich/govalidator"
	"errors"
)

type TransactionRepositoryInterface interface {
	RegisterKey(transaction *Transaction) error
	Save(transaction *Transaction) error
	Find(id string) (*Transaction, error)
}

const (
	TransactionPending string = "pending"
	TransactionCompleted string = "completed"
	TransactionError string = "error"
	TransactionConfirmed string = "confirmed"
)

type Transactions struct {
	Transaction []Transaction
}

type Trasaction struct {
	Base												`valid: "required"`
	AccountFrom				*Account 	`valid: "-"`
	Amount						float64 	`json: "amount" valid: "notnull"`
	PixKeyTo					*PixKey 	`valid: "-"`
	Status						string 		`json: "status" valid: "notnull"`
	Description				string 		`json: "description" valid: "notnull"`
	CancelDescription	string 		`json: "cancelDescription" valid: "-"`
}

func (transaction *Trasaction) isValid() error {
	_, err := govalidator.ValidateStruct(transaction)

	if transaction.Amount <=0 {
		return errors.New("the amount must be greater than 0")
	}

	if transaction.Status != TransactionPending && transaction.Status != TransactionCompleted && transaction.Status != TransactionError {
		return errors.New("invalid status for the transaction")
	}

	if transaction.PixKeyTo.AccountId == transaction.AccountFrom.ID {
		return errors.New("the source and destination accounts cannot be the same")
	}

	if err != nil {
		return err
	}
	return nil
}

func NewTrasaction(accountFrom *Account, amount float64, pixKeyTo *PixKey, description string) (*Trasaction, error) {
	trasaction := Trasaction {
		AccountFrom: accountFrom,
		Amount: amount,
		PixKeyTo: pixKeyTo,
		Status: TransactionPending,
		Description: description,
	}

	trasaction.ID = uuid.NewV4().String()
	trasaction.CreatedAt = time.Now()

	err := trasaction.isValid()

	if err != nil {
		return nil, err
	}

	return &trasaction, nil
}

func (transaction *Transaction) Complete() error {
	transaction.Status = TransactionCompleted
	transaction.UpdatedAt = time.Now()

	err := transaction.isValid()

	return err
}

func (transaction *Transaction) Confirm() error {
	transaction.Status = TransactionConfirmed
	transaction.UpdatedAt = time.Now()

	err := transaction.isValid()

	return err
}

func (transaction *Transaction) Cancel(description string) error {
	transaction.Status = TransactionError
	transaction.UpdatedAt = time.Now()
	transaction.Description = description

	err := transaction.isValid()

	return err
}