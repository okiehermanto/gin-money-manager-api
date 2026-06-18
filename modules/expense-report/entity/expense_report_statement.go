package entity

import (
	"gin-money-manager-api/modules/shared/money"
	userentity "gin-money-manager-api/modules/user/entity"
	"time"

	"github.com/google/uuid"
)

type ExpenseReportStatement struct {
	ID              uuid.UUID       `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	ExpenseReportID uuid.UUID       `json:"expense_report_id"`
	ExpenseReport   ExpenseReport   `gorm:"foreignKey:ExpenseReportID;referenced:ID"`
	AuthorID        uuid.UUID       `json:"author_id"`
	Author          userentity.User `gorm:"foreignKey:AuthorId;references:ID"`
	Name            string          `json:"name"`
	Description     string          `json:"description"`
	Amount          money.Money     `json:"amount" gorm:"embedded;embeddedPrefix:amount_"`
	Settled         bool            `json:"settled"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
}
