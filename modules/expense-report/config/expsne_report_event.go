package config

import (
	"gin-money-manager-api/modules/expense-report/event"
	"gin-money-manager-api/modules/expense-report/listener"
	eventmanager "gin-money-manager-api/modules/shared/event_manager"
)

func RegisterEvents(container *ExpenseReportContainer) {
	calculateExpenseReportBalance := listener.NewCalculateExpenseReportBalance(
		container.ExpenseReportRepository,
		container.ExpenseReportStatementRepository,
	)

	eventmanager.Dispatcher.On(event.ExpenseReportStatementCreatedEvent{}.Name(), func(e eventmanager.Event) {
		calculateExpenseReportBalance.Handler(e)
	})
}
