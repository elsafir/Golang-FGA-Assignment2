package models

import (
	"time"
)

type Item struct {
	ItemId      uint      `json:"itemId" gorm:"primary_key"`
	ItemCode    string    `json:"itemCode"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	OrderId     uint      `json:"orderId"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
