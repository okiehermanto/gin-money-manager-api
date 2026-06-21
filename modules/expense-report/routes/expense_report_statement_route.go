package routes

import (
	"gin-money-manager-api/modules/auth/middleware"
	"gin-money-manager-api/modules/expense-report/command"
	"gin-money-manager-api/modules/expense-report/config"
	"gin-money-manager-api/modules/expense-report/query"
	"gin-money-manager-api/modules/expense-report/resource"

	"github.com/gin-gonic/gin"
)

func RegisterExpenseReportStatementRoute(r *gin.Engine, container *config.ExpenseReportContainer) {
	auth := middleware.Auth(container.UserRepository)

	createExpenseReportStatement := command.NewCreateExpenseReportStatement(
		container.ExpenseReportStatementRepository,
		container.ExpenseReportRepository,
	)

	GetAttachmentExpenseReportstatment := query.NewGetAttachmentExpenseReportStatement(
		container.ExpenseReportStatementRepository,
	)

	expenseReportExpenseReportStatementResource := resource.NewExpenseReportExpenseReportStatementResource(container.ExpenseReportStatementRepository)

	expenseReportStatement := r.Group("/expense-reports/:expense_report/statements")

	expenseReportStatement.POST(
		"",
		auth,
		createExpenseReportStatement.Handler,
	)

	expenseReportStatement.GET(
		"",
		auth,
		expenseReportExpenseReportStatementResource.Index,
	)
	expenseReportStatement.GET(
		"/:expense_report_statement",
		auth,
		expenseReportExpenseReportStatementResource.Index,
	)

	expenseReportStatement.GET(
		"/:expense_report_statement/attachment",
		auth,
		GetAttachmentExpenseReportstatment.Handler,
	)
}
