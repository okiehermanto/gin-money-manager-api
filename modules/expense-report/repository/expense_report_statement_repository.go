package repository

import (
	"gin-money-manager-api/modules/expense-report/entity"
	"gin-money-manager-api/modules/shared/repository"

	"gorm.io/gorm"
)

type expenseReportStatementRepository struct {
	repository.BaseRepository[entity.ExpenseReportStatement]
}

func NewExpenseReportStatementRepository(db *gorm.DB) ExpenseReportStatementRepository {
	return &expenseReportStatementRepository{
		BaseRepository: repository.NewBaseRepository[entity.ExpenseReportStatement](db),
	}
}
