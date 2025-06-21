package dtos

type LoginRequest struct {
	Username string
	Password string
}

type LoginResponse struct {
	Token   string
	Cookies string
}

type RegisterRequest struct {
	Username string
	Password string
}
