package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID        uuid.UUID  `gorm:"column:id"`
	SKU       string     `gorm:"column:sku"`
	Name      string     `gorm:"column:name"`
	CreatedBy uuid.UUID  `gorm:"column:createdby"`
	CreatedAt time.Time  `gorm:"column:createdat"`
	Status    string     `gorm:"column:status"`
	UpdatedBy uuid.UUID  `gorm:"column:updatedby"`
	UpdatedAt *time.Time `gorm:"column:updatedat;autoUpdateTime:false"`
}

func (prd Product) TableName() string {
	return "dataproduct.masterproduct"
}
