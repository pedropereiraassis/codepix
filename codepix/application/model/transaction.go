package model

import (
	"encoding/json"
	"fmt"

	"github.com/asaskevich/govalidator"
)

type Transaction struct {
	ID           string  `json:"id" validate:"required,uuid"`
	AccountId    string  `json:"accountId" validate:"required,uuid"`
	Amount       float64 `json:"amount" validate:"required,numeric"`
	PixKeyTo     string  `json:"pixKeyTo" validate:"required"`
	PixKeyKindTo string  `json:"pixKeyKindTo" validate:"required"`
	Description  string  `json:"description" validate:"required"`
	Status       string  `json:"status" validate:"required"`
	Error        string  `json:"error"`
}

func (t *Transaction) isValid() error {
	_, err := govalidator.ValidateStruct(t)

	if err != nil {
		return fmt.Errorf("error during transaction validation: %s", err.Error())
	}

	return nil
}

func (t *Transaction) ParseJson(data []byte) error {
	err := json.Unmarshal(data, t)
	if err != nil {
		return err
	}

	err = t.isValid()
	if err != nil {
		return err
	}

	return nil
}

func (t *Transaction) ToJson() ([]byte, error) {
	err := t.isValid()
	if err != nil {
		return nil, err
	}

	result, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func NewTransaction() *Transaction {
	return &Transaction{}
}
