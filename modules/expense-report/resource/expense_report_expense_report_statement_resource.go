package resource

import (
	"gin-money-manager-api/modules/expense-report/entity"
	"gin-money-manager-api/modules/expense-report/repository"
	"gin-money-manager-api/modules/shared/dto"
	"gin-money-manager-api/modules/shared/resource"
)

type ExpenseReportExpenseReportStatementResource struct {
	*resource.BaseResource[
		entity.ExpenseReportStatement,
		dto.EmptyDto,
		dto.EmptyDto,
	]
	ExpenseReport string
}

func NewExpenseReportExpenseReportStatementResource(
	repository repository.ExpenseReportStatementRepository,
) *ExpenseReportExpenseReportStatementResource {
	return &ExpenseReportExpenseReportStatementResource{
		BaseResource: &resource.BaseResource[entity.ExpenseReportStatement, dto.EmptyDto, dto.EmptyDto]{
			Repository: repository,
			RouteParam: entity.ExpenseReportStatement{}.EntityName(),

			Filter: "expense_report",
		},
	}
}
