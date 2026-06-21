package listener

import (
	"fmt"
	"gin-money-manager-api/modules/expense-report/entity"
	"gin-money-manager-api/modules/expense-report/repository"
	valueobject "gin-money-manager-api/modules/expense-report/value-object"
	eventmanager "gin-money-manager-api/modules/shared/event_manager"
)

type CalculateExpenseReportBalance struct {
	repository                       repository.ExpenseReportRepository
	expenseReportStatementRepository repository.ExpenseReportStatementRepository
}

func NewCalculateExpenseReportBalance(
	repository repository.ExpenseReportRepository,
	expenseReportStatementRepository repository.ExpenseReportStatementRepository,
) *CalculateExpenseReportBalance {
	return &CalculateExpenseReportBalance{
		repository:                       repository,
		expenseReportStatementRepository: expenseReportStatementRepository,
	}
}

func (h *CalculateExpenseReportBalance) Handler(event eventmanager.Event) {
	if statement, ok := event.Data().(*entity.ExpenseReportStatement); ok {
		switch statement.ExpenseReport.Type {
		case valueobject.ExpenseReportTypeReport:
			statement.Settled = true
			expenseReport := statement.ExpenseReport

			newBalance, err := expenseReport.Balance.Add(statement.Amount)

			if err != nil {
				fmt.Print("todo: add failed jobs")
			}

			expenseReport.Balance = newBalance

			h.expenseReportStatementRepository.Update(*statement)
			h.repository.Update(*expenseReport)
		default:
			return
		}
	}

}
