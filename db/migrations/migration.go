package migrations

import (
	"github.com/dmgol/mashunya/app/models"
	"github.com/dmgol/mashunya/db"
	"github.com/qor/media_library"
)

func init() {
	AutoMigrate(&media_library.AssetManager{})

	AutoMigrate(&models.Product{}, &models.ProductImage{}, &models.ColorVariation{}, &models.ColorVariationImage{}, &models.SizeVariation{})
	AutoMigrate(&models.Color{}, &models.Size{}, &models.Category{}, &models.Collection{})

	AutoMigrate(&models.Address{})

	AutoMigrate(&models.Order{}, &models.OrderItem{})

	AutoMigrate(&models.User{})

	AutoMigrate(&models.MediaLibrary{})
}

func AutoMigrate(values ...interface{}) {
	for _, value := range values {
		db.DB.AutoMigrate(value)
	}
}
