package handlers

import (
	"ghgist-blog/models"
	"ghgist-blog/utils"

	"github.com/gin-gonic/gin"
)

// CreateTokenForUser creates a JWT token for the given user
// Returns token string and error
func CreateTokenForUser(user *models.User) (string, error) {
	return utils.CreateToken(user.ID.Hex(), user.Role)
}

// SendAuthResponse sends a standardized authentication response
func SendAuthResponse(c *gin.Context, message string, data interface{}, statusCode int) {
	response := models.JSONresponse{
		Message: message,
		Data:    data,
	}
	c.JSON(statusCode, response)
}
