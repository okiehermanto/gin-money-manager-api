package expensereport

import (
	"gin-money-manager-api/modules/expense-report/container"
	"gin-money-manager-api/modules/expense-report/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(r *gin.Engine, db *gorm.DB) {
	container := container.NewExpenseReportContainer(db)

	routes.RegisterExpenseReportRoute(r, container)
}
