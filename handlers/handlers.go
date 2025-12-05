package handlers

import (
	"ghgist-blog/models"
	"ghgist-blog/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ---------------------------Structs----------------------------
type Handler struct {
	RegisterService services.RegisterInterface
	LoginService    services.LoginInterface
	FetchService    services.FetchWritersInterface
}

func NewHandlers(services *ServiceContainer) *Handler {
	return &Handler{
		RegisterService: services.RegisterService,
		LoginService:    services.LoginService,
		FetchService:    services.FetchService,
	}
}

//---------------------------Handler functions----------------------------

func (h *Handler) Register(c *gin.Context) {
	// Validate HTTP method

	var req models.User
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	// Call service - let it handle all business logic
	user, err := h.RegisterService.RegisterUser(req.Email, req.Username, req.Password, req.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User registration failed: " + err.Error()})
		return
	}
	// Create token
	token, err := CreateTokenForUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token creation failed: " + err.Error()})
		return
	}
	data := map[string]interface{}{
		"token": token,
		"user":  user,
	}

	// Success response
	SendAuthResponse(c, "User created successfully", data, http.StatusCreated)

}

func (h *Handler) Login(c *gin.Context) {

	var req models.User
	//decoding the json request
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload" + err.Error()})
		return
	}

	user, err := h.LoginService.LoginUser(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Login failed: " + err.Error()})
		return
	}
	// Create token
	token, err := CreateTokenForUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token creation failed: " + err.Error()})
		return
	}
	data := map[string]interface{}{
		"token": token,
		"user":  user,
	}
	// Success response
	SendAuthResponse(c, "Login successful", data, http.StatusCreated)
}
func (h *Handler) FetchAllWriters(c *gin.Context) {

	users, err := h.FetchService.FetchAllWriters()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch users" + err.Error()})
		return
	}
	data := map[string]interface{}{
		"users": users,
	}
	//success response
	SendAuthResponse(c, "Fetch successful", data, http.StatusCreated)
}
