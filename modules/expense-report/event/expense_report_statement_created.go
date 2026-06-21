package event

import "gin-money-manager-api/modules/expense-report/entity"

type ExpenseReportStatementCreatedEvent struct {
	ExpenseReportStatement *entity.ExpenseReportStatement
}

func (e ExpenseReportStatementCreatedEvent) Name() string {
	return "expense_report_statement_created"
}

func (e ExpenseReportStatementCreatedEvent) Data() interface{} {
	return e.ExpenseReportStatement
}
