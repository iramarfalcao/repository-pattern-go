package migrations

import (
	"github.com/iramarfalcao/repository-pattern-go/models/product"
	"gorm.io/gorm"
)

func RunAutoMigrations(db *gorm.DB) {
	db.AutoMigrate(&product.Product{})
}
