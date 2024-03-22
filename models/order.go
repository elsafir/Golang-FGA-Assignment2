package models

import (
	"time"
)

type Order struct {
	OrderId      uint      `json:"orderId" gorm:"primary_key"`
	CustomerName string    `json:"customerName"`
	OrderedAt    string    `json:"orderedAt"`
	Items        []Item    `json:"items" gorm:"foreignKey:OrderId"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
