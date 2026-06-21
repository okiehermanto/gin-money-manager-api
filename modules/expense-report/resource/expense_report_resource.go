package resource

import (
	"gin-money-manager-api/modules/expense-report/entity"
	"gin-money-manager-api/modules/expense-report/repository"
	"gin-money-manager-api/modules/shared/dto"
	"gin-money-manager-api/modules/shared/resource"
)

type ExpenseReportResource struct {
	*resource.BaseResource[
		entity.ExpenseReport,
		dto.EmptyDto,
		dto.EmptyDto,
	]
}

func NewExpenseReportResource(repository repository.ExpenseReportRepository) *ExpenseReportResource {
	return &ExpenseReportResource{
		BaseResource: &resource.BaseResource[entity.ExpenseReport, dto.EmptyDto, dto.EmptyDto]{
			Repository: repository,
			RouteParam: entity.ExpenseReport{}.EntityName(),
			Relationships: []string{
				"User.Roles",
			},
			Paginate: true,
		},
	}
}
