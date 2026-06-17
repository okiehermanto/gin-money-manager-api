package repository

import (
	"gin-money-manager-api/modules/expense-report/entity"
	"gin-money-manager-api/modules/shared/repository"
)

type ExpenseReportRepository interface {
	repository.BaseRepository[entity.ExpenseReport]
}
