package dto

type CreateTenantRequest struct {
	Name          string `json:"name" binding:"required"`
	Plan          string `json:"plan" binding:"required"`
	AdminEmail    string `json:"admin_email" binding:"required,email"`
	AdminPassword string `json:"admin_password" binding:"required,min=6"`
}
