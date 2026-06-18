package command

import (
	"gin-money-manager-api/modules/expense-report/dto"
	"gin-money-manager-api/modules/expense-report/entity"
	userentity "gin-money-manager-api/modules/user/entity"

	"gin-money-manager-api/modules/expense-report/repository"
	valueobject "gin-money-manager-api/modules/expense-report/value-object"
	"gin-money-manager-api/modules/shared/helper"
	"gin-money-manager-api/modules/shared/money"
	"gin-money-manager-api/modules/shared/response"

	"github.com/gin-gonic/gin"
)

type CreateExpenseReport struct {
	repository repository.ExpenseReportRepository
}

func NewCreatExpenseReport(repository repository.ExpenseReportRepository) *CreateExpenseReport {
	return &CreateExpenseReport{
		repository: repository,
	}
}

func (h *CreateExpenseReport) Handler(c *gin.Context) {
	var body dto.CreateExpenseReport

	if !helper.BindAndValidate(c, &body) {
		return
	}

	user := c.MustGet("user").(userentity.User)

	expenseReport := entity.ExpenseReport{
		User: user,
		Name: body.Name,
		Type: valueobject.ExpenseReportType(body.Type),
		Balance: money.Money{
			Amount:   0,
			Currency: "IDR",
		},
	}

	expenseReport, err := h.repository.Create(expenseReport)

	response.Success(c, expenseReport, err)
}
