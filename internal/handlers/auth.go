package handlers

import (
	"encoding/json"
	"fmt"
	"miras/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp() gin.HandlerFunc {
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

func (h *Handler) signIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user *models.Login

		err := json.NewDecoder(c.Request.Body).Decode(&user)
		if err != nil {
			
			c.JSON(400, gin.H{"error": fmt.Sprintf("Bad request (%s)", err.Error())})
			return
		}

		cookie, err := h.Service.SigninService(user)
		if err != nil {
			c.JSON(400, gin.H{"error": fmt.Sprintf("Bad request (%s)", err.Error())})
			return
		}

		c.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
		c.JSON(200, gin.H{"Status": "User successfuly signed"})

	}
}
