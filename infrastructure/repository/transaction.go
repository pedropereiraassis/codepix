package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/pedropereiraassis/codepix/domain/model"
)

type TransactionRepositoryDB struct {
	Db *gorm.DB
}

func (r TransactionRepositoryDB) Register(transaction *model.Transaction) error {
	err := r.Db.Create(transaction).Error

	if err != nil {
		return err
	}

	return nil
}

func (r TransactionRepositoryDB) Save(transaction *model.Transaction) error {
	err := r.Db.Save(transaction).Error

	if err != nil {
		return err
	}

	return nil
}

func (r TransactionRepositoryDB) Find(id string) (*model.Transaction, error) {
	var transaction model.Transaction

	r.Db.Preload("AccountFrom.Bank").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, fmt.Errorf("no transaction was found")
	}

	return &transaction, nil
}