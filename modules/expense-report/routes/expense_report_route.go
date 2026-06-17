package routes

import (
	"gin-money-manager-api/modules/auth/middleware"
	"gin-money-manager-api/modules/expense-report/command"
	"gin-money-manager-api/modules/expense-report/container"

	"github.com/gin-gonic/gin"
)

func RegisterExpenseReportRoute(r *gin.Engine, container *container.ExpenseReportContainer) {
	createExpenseReport := command.NewCreatExpenseReport(container.ExpenseReportRepository)

	expenseReports := r.Group("/expense-reports")
	expenseReports.Use(
		middleware.Auth(container.UserRepository),
		middleware.HasRoles("user"),
	)
	{
		expenseReports.POST("", createExpenseReport.Handler)
	}
}
