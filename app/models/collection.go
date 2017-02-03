package models

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/l10n"
)

type Collection struct {
	gorm.Model
	Name string
	l10n.LocaleCreatable
	Products []Product `gorm:"many2many:product_collections;"`
}
