package config

import (
	"gin-money-manager-api/modules/expense-report/repository"
	userrepository "gin-money-manager-api/modules/user/repository"

	"gorm.io/gorm"
)

type ExpenseReportContainer struct {
	UserRepository                   userrepository.UserRepository
	ExpenseReportRepository          repository.ExpenseReportRepository
	ExpenseReportStatementRepository repository.ExpenseReportStatementRepository
}

func NewExpenseReportContainer(db *gorm.DB) *ExpenseReportContainer {
	userRepository := userrepository.NewUserRepository(db)
	expenseReportRepository := repository.NewExpenseReportRepository(db)
	expenseReportStatementRepository := repository.NewExpenseReportStatementRepository(db)

	return &ExpenseReportContainer{
		UserRepository:                   userRepository,
		ExpenseReportRepository:          expenseReportRepository,
		ExpenseReportStatementRepository: expenseReportStatementRepository,
	}
}
