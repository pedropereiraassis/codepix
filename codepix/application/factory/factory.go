package factory

import (
	"github.com/jinzhu/gorm"
	"github.com/pedropereiraassis/codepix/application/usecase"
	"github.com/pedropereiraassis/codepix/infrastructure/repository"
)

func TransactionUseCaseFactory(database *gorm.DB) usecase.TransactionUseCase {
	pixRepository := repository.PixKeyRepositoryDB{Db: database}
	transactionRepository := repository.TransactionRepositoryDB{Db: database}

	transactionUseCase := usecase.TransactionUseCase{
		TransactionRepository: &transactionRepository,
		PixRepository: pixRepository,
	}

	return transactionUseCase
}