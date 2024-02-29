package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) editUserData() gin.HandlerFunc {

	return func(c *gin.Context) {

		cookie, err := c.Request.Cookie("Token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Unauthorized (%s)", err.Error())})
			return
		}

		user, err := h.Service.Edit.GetUserByToken(cookie.Value)
		if err != nil {
			c.JSON(400, gin.H{"error": fmt.Sprintf("Bad request (%s)", err.Error())})
			return
		}
		
		

	}
}
