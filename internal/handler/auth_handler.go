package handler

import (
	"cbt/internal/models"
	"cbt/internal/service"
	"cbt/internal/utils"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthenticationHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthenticationHandler {
	return &AuthenticationHandler{authService: authService}
}

func (h *AuthenticationHandler) Register(c *gin.Context) {
	var req models.RegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	user, err := h.authService.RegisterUser(req)
	if err != nil {
		if err.Error() == "username already exists" || err.Error() == "role not found: "+req.RoleName {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user: " + err.Error()})
		}
		return
	}

	response := models.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		FullName:  user.FullName,
		Email:     user.Email,
		RoleName:  user.Role.RoleName,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	c.JSON(http.StatusCreated, response)
}

func (h *AuthenticationHandler) Login(c *gin.Context) {
	expirationTime := time.Now().Add(86400 * time.Second)
	var req models.LoginUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	user, token, err := h.authService.LoginUser(req)
	if err != nil {
		if errors.Is(err, errors.New("invalid username or password")) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Login failed: " + err.Error()})
		}
		return
	}

	encodedCookies, _ := utils.SetEncryptCookies(token)

	loginResp := models.LoginResponse{
		User: models.UserResponse{
			ID:       user.ID,
			Username: user.Username,
			FullName: user.FullName,
			Email:    user.Email,
			RoleName: user.Role.RoleName,
		},
		Token: token,
	}
	cookieValue := fmt.Sprintf("%d|%s|%d", expirationTime.Unix(), encodedCookies, 86400)
	user_id := user.Base.ID
	c.SetCookie("session_token", cookieValue, 86400, "/", "https://cbt.edunusa.co.id", true, true)
	c.Header("Session-UUID", user_id.String())
	c.JSON(http.StatusOK, loginResp)
}

func (h *AuthenticationHandler) Logout(c *gin.Context) {
	c.SetCookie("session_token", "", -1, "/", "https://cbt.edunusa.co.id", true, true) // Menghapus cookie
	c.JSON(http.StatusOK, gin.H{"message": "Logout berhasil!"})
}
