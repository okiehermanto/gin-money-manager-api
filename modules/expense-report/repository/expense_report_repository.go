package repository

import (
	"gin-money-manager-api/modules/expense-report/entity"
	"gin-money-manager-api/modules/shared/repository"

	"gorm.io/gorm"
)

type expenseReportRepository struct {
	repository.BaseRepository[entity.ExpenseReport]
}

func NewExpenseReportRepository(db *gorm.DB) ExpenseReportRepository {
	return &expenseReportRepository{
		BaseRepository: repository.NewBaseRepository[entity.ExpenseReport](db),
	}
}
