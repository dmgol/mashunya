package auth

import (
	"log"

	"github.com/dmgol/mashunya/app/models"
	"github.com/dmgol/mashunya/db"
	"gopkg.in/authboss.v0"
)

type AuthStorer struct {
}

func (s AuthStorer) Create(key string, attr authboss.Attributes) error {
	log.Println("AuthStorer.Create ...", key)
	var user models.AdminUser
	if err := attr.Bind(&user, true); err != nil {
		return err
	}

	if err := db.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (s AuthStorer) Put(key string, attr authboss.Attributes) error {
	log.Println("AuthStorer.Put ...", key)
	var user models.AdminUser
	if err := db.DB.Where("email = ?", key).First(&user).Error; err != nil {
		return authboss.ErrUserNotFound
	}

	if err := attr.Bind(&user, true); err != nil {
		return err
	}

	if err := db.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (s AuthStorer) Get(key string) (result interface{}, err error) {
	log.Println("AuthStorer.Get ...", key)
	var user models.AdminUser
	if err := db.DB.Where("email = ?", key).First(&user).Error; err != nil {
		return nil, authboss.ErrUserNotFound
	}
	log.Println("... AuthStorer.Get", user)
	return &user, nil
}

func (s AuthStorer) ConfirmUser(tok string) (result interface{}, err error) {
	log.Println("AuthStorer.ConfirmUser ...", tok)
	var user models.AdminUser
	if err := db.DB.Where("confirm_token = ?", tok).First(&user).Error; err != nil {
		return nil, authboss.ErrUserNotFound
	}
	return &user, nil

	return nil, authboss.ErrUserNotFound
}

func (s AuthStorer) RecoverUser(rec string) (result interface{}, err error) {
	log.Println("AuthStorer.ConfirmUser ...", rec)
	var user models.AdminUser
	if err := db.DB.Where("recover_token = ?", rec).First(&user).Error; err != nil {
		return nil, authboss.ErrUserNotFound
	}
	return &user, nil

	return nil, authboss.ErrUserNotFound
}
