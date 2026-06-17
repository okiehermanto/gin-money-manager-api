package entity

import (
	expensereportvalueobject "gin-money-manager-api/modules/expense-report/value-object"
	valueobject "gin-money-manager-api/modules/shared/value-object"

	"time"

	"github.com/google/uuid"
)

type ExpenseReport struct {
	ID        uuid.UUID                                  `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    uuid.UUID                                  `json:"user_id"`
	Name      string                                     `json:"name"`
	Type      expensereportvalueobject.ExpenseReportType `json:"type"`
	Balance   valueobject.Money                          `json:"balance" gorm:"embedded;embeddedPrefix:balance_"`
	CreatedAt time.Time                                  `json:"created_at"`
	UpdatedAt time.Time                                  `json:"updated_at"`
}
