package expensereport

import (
	"gin-money-manager-api/modules/expense-report/config"
	"gin-money-manager-api/modules/expense-report/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(r *gin.Engine, db *gorm.DB) {
	container := config.NewExpenseReportContainer(db)

	config.RegisterEvents(container)

	routes.RegisterExpenseReportStatementRoute(r, container)
	routes.RegisterExpenseReportRoute(r, container)

}
