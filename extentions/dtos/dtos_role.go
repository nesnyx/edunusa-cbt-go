package dtos

type RoleRequest struct {
	ID   string `json:"id"`
	Name string `json:"name" binding:"required"`
}

type RoleResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
