package command

import (
	"fmt"
	"gin-money-manager-api/modules/expense-report/dto"
	"gin-money-manager-api/modules/expense-report/entity"
	"gin-money-manager-api/modules/expense-report/repository"
	"gin-money-manager-api/modules/shared/helper"
	"gin-money-manager-api/modules/shared/money"
	"gin-money-manager-api/modules/shared/response"
	userentity "gin-money-manager-api/modules/user/entity"

	"github.com/gin-gonic/gin"
)

type CreateExpenseReportStatement struct {
	repository              repository.ExpenseReportStatementRepository
	expenseReportRepository repository.ExpenseReportRepository
}

func NewCreateExpenseReportStatement(
	repository repository.ExpenseReportStatementRepository,
	expenseReportRepository repository.ExpenseReportRepository,
) *CreateExpenseReportStatement {
	return &CreateExpenseReportStatement{
		repository:              repository,
		expenseReportRepository: expenseReportRepository,
	}
}

func (h *CreateExpenseReportStatement) Handler(c *gin.Context) {
	var body dto.CreateExpenseReportStatement
	if !helper.BindAndValidate(c, &body) {
		return
	}

	expenseReport, err := h.expenseReportRepository.Find(c.Param("expense_report"), nil)

	if err != nil {
		response.ModelNotFound(c, expenseReport.EntityName())
		return
	}

	user := c.MustGet("user").(userentity.User)

	attachment, err := helper.SaveAttachment(
		fmt.Sprintf("storage/expense-reports/%s/attachments/", expenseReport.ID),
		body.Attachment)

	expenseReportStatement := entity.ExpenseReportStatement{
		ExpenseReport: &expenseReport,
		Author:        &user,
		Name:          body.Name,
		Description:   body.Description,
		Attachment:    attachment,
		Amount: money.Money{
			Amount:   int64(body.Amount),
			Currency: "IDR",
		},
	}

	expenseReportStatement, saveError := h.repository.Create(expenseReportStatement)

	response.Success(c, expenseReportStatement, saveError)

}
