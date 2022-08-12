package admin

import "github.com/motikingo/websocketproject/internal/pkg/entity"

// AdminService interface
type AdminService interface {
	CreateAdmin(admin *entity.Admin) *entity.Admin
	GetAdminByID(id string ) *entity.Admin
	DeleteAdminByID( id string  ) bool 
	GetAdminByEmail(email string )  *entity.Admin 
	SaveAdmin(admin *entity.Admin)  *entity.Admin 
	AdminEmailExist(   email string ) bool 
}