package resource

import (
	"gin-money-manager-api/modules/expense-report/entity"
	"gin-money-manager-api/modules/expense-report/repository"
	"gin-money-manager-api/modules/shared/dto"
	"gin-money-manager-api/modules/shared/resource"
)

type UserExpenseReportResource struct {
	*resource.BaseResource[
		entity.ExpenseReport,
		dto.EmptyDto,
		dto.EmptyDto,
	]
	UserID string
}

func NewUserExpenseReportResource(repository repository.ExpenseReportRepository) *ExpenseReportResource {
	return &ExpenseReportResource{
		BaseResource: &resource.BaseResource[entity.ExpenseReport, dto.EmptyDto, dto.EmptyDto]{
			Repository: repository,
			Relationships: []string{
				"User.Roles",
			},
			Filter: "user_id",
		},
	}
}
