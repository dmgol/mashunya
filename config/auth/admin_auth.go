package auth

import (
	"log"
	"net/http"

	"github.com/dmgol/mashunya/app/models"
	"github.com/qor/admin"
	"github.com/qor/qor"
	"github.com/qor/roles"
)

func init() {
	roles.Register("admin", func(req *http.Request, currentUser interface{}) bool {
		return currentUser != nil && currentUser.(*models.AdminUser).Role == "Admin"
	})
}

type AdminAuth struct {
}

func (AdminAuth) LoginURL(c *admin.Context) string {
	return "/auth/login"
}

func (AdminAuth) LogoutURL(c *admin.Context) string {
	return "/auth/logout"
}

func (AdminAuth) GetCurrentUser(c *admin.Context) qor.CurrentUser {
	log.Println("AdminAuth.GetCurrentUser...")
	userInter, err := Auth.CurrentUser(c.Writer, c.Request)
	if userInter != nil && err == nil {
		log.Println("...AdminAuth.GetCurrentUser", userInter)
		return userInter.(*models.AdminUser)
	}
	log.Println("...AdminAuth.GetCurrentUser", nil)
	return nil
}
