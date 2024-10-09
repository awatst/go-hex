package adapters

import (
	"github.com/awatst/go-hex/core"
	"gorm.io/gorm"
)

// secondary adapter
type GormOrderRepository struct {
	db *gorm.DB
}

func NewGormOrderRepository(db *gorm.DB) core.OrderRepository {
	return &GormOrderRepository{db: db}
}

// func save to db
func (r *GormOrderRepository) Save(order core.Order) error {
	if result := r.db.Create(&order); result.Error != nil {
		return result.Error
	}

	return nil
}
