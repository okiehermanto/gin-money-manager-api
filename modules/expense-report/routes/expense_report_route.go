package routes

import (
	"gin-money-manager-api/modules/auth/middleware"
	"gin-money-manager-api/modules/expense-report/command"
	"gin-money-manager-api/modules/expense-report/container"
	"gin-money-manager-api/modules/expense-report/resource"

	"github.com/gin-gonic/gin"
)

func RegisterExpenseReportRoute(r *gin.Engine, container *container.ExpenseReportContainer) {
	createExpenseReport := command.NewCreatExpenseReport(container.ExpenseReportRepository)

	ExpenseReportResource := resource.NewExpenseReportResource(container.ExpenseReportRepository)
	userExpenseReportResource := resource.NewUserExpenseReportResource(container.ExpenseReportRepository)

	userExpenseReport := r.Group("/users/:user_id/expense-reports")

	userExpenseReport.GET("",
		userExpenseReportResource.Index,
	)

	expenseReports := r.Group("/expense-reports")

	expenseReports.POST("",
		middleware.Auth(container.UserRepository),
		middleware.HasRoles("admin"),
		createExpenseReport.Handler,
	)

	expenseReports.GET("",
		middleware.Auth(container.UserRepository),
		middleware.HasRoles("admin"),
		ExpenseReportResource.Index,
	)

	expenseReports.GET("/:id",
		middleware.Auth(container.UserRepository),
		middleware.HasRoles("admin"),
		//anothermiddleware
		ExpenseReportResource.Index,
	)
}
