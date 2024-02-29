package handlers

import (
	"encoding/json"
	"fmt"
	"miras/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user *models.Register

		err := json.NewDecoder(c.Request.Body).Decode(&user)
		if err != nil {
			c.JSON(400, gin.H{"error": "Bad request"})
			return
		}

		err = h.Service.Auth.SignupService(user)
		if err != nil {
			c.JSON(400, gin.H{"error": fmt.Sprintf("Bad request (%s)", err.Error())})
			return
		}

		c.JSON(200, gin.H{"Status": "User successfuly create"})
	}
}
