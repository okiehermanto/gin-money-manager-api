package dto

import (
	"mime/multipart"
)

type CreateExpenseReportStatement struct {
	Name        string                `form:"name" binding:"required"`
	Description string                `form:"description" binding:"required"`
	Attachment  *multipart.FileHeader `form:"attachment" binding:"required"`
	Amount      int                   `form:"amount" binding:"required"`
}
