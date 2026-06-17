package dto

import valueobject "gin-money-manager-api/modules/expense-report/value-object"

type CreateExpenseReport struct {
	Name string                        `json:"name" binding:"required"`
	Type valueobject.ExpenseReportType `json:"type" binding:"required,oneof=reimburse report"`
}
