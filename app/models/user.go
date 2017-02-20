package models

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email      string
	Password   string
	FirstName  sql.NullString
	MiddleName sql.NullString
	LastName   sql.NullString
	Wholesaler bool
	Gender     string
	Addresses  []Address

	// Confirm
	ConfirmToken string
	Confirmed    bool

	// Recover
	RecoverToken       string
	RecoverTokenExpiry *time.Time
}

func (user User) DisplayName() string {
	return user.Email
}

func (user User) AvailableLocales() []string {
	return []string{"en-US", "ru-RU"}
}

type AdminUser struct {
	gorm.Model
	Email    string
	Password string
	Role     string

	// Confirm
	ConfirmToken string
	Confirmed    bool

	// Recover
	RecoverToken       string
	RecoverTokenExpiry *time.Time
}

func (user AdminUser) DisplayName() string {
	return user.Email
}

func (user AdminUser) AvailableLocales() []string {
	return []string{"en-US", "ru-RU"}
}
