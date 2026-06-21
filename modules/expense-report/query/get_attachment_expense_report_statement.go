package query

import (
	"gin-money-manager-api/modules/expense-report/repository"
	"gin-money-manager-api/modules/shared/response"

	"github.com/gin-gonic/gin"
)

type GetAttachmentExpenseReportstatment struct {
	repository repository.ExpenseReportStatementRepository
}

func NewGetAttachmentExpenseReportStatement(
	repository repository.ExpenseReportStatementRepository,
) *GetAttachmentExpenseReportstatment {
	return &GetAttachmentExpenseReportstatment{
		repository: repository,
	}
}

func (h *GetAttachmentExpenseReportstatment) Handler(c *gin.Context) {
	expenseReportStatementId := c.Param("expense_report_statement")

	expenseReportStatement, err := h.repository.Find(expenseReportStatementId, nil)

	if err != nil {
		response.ModelNotFound(c, expenseReportStatement.EntityName())
		return
	}

	c.File(expenseReportStatement.Attachment)
}
