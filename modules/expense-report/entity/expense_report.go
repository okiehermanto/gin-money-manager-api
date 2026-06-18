package entity

import (
	expensereportvalueobject "gin-money-manager-api/modules/expense-report/value-object"
	"gin-money-manager-api/modules/shared/money"
	userentity "gin-money-manager-api/modules/user/entity"

	"time"

	"github.com/google/uuid"
)

type ExpenseReport struct {
	ID        uuid.UUID                                  `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    uuid.UUID                                  `json:"user_id"`
	User      userentity.User                            `json:"user" gorm:"foreignKey:UserID;references:ID"`
	Name      string                                     `json:"name"`
	Type      expensereportvalueobject.ExpenseReportType `json:"type"`
	Balance   money.Money                                `json:"balance" gorm:"embedded;embeddedPrefix:balance_"`
	CreatedAt time.Time                                  `json:"created_at"`
	UpdatedAt time.Time                                  `json:"updated_at"`
}
