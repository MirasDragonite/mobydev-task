package handlers

import (
	"encoding/json"
	"fmt"
	"miras/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func errorHandler(c *gin.Context, code int, err error) {
	c.JSON(code, gin.H{"error": fmt.Sprintf("Bad Request (%s)", err.Error())})
}

func (h *Handler) getUserDataForEdit() gin.HandlerFunc {

	return func(c *gin.Context) {

		idStr := c.Request.URL.Query().Get("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			errorHandler(c, 400, err)
			return

		}
		cookie, err := c.Request.Cookie("Token")
		if err != nil {
			errorHandler(c, 401, err)
			return
		}

		user, err := h.Service.Edit.GetUserData(cookie.Value, id)

		if err != nil {
			errorHandler(c, 400, err)
			return
		}

		c.JSON(200, user)
	}

}

func (h *Handler) editUserData() gin.HandlerFunc {

	return func(c *gin.Context) {

		idStr := c.Request.URL.Query().Get("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			errorHandler(c, 400, err)
			return

		}
		cookie, err := c.Request.Cookie("Token")
		if err != nil {
			errorHandler(c, 401, err)
			return
		}

		var userEdit models.UserEdit

		err = json.NewDecoder(c.Request.Body).Decode(&userEdit)
		if err != nil {
			errorHandler(c, 400, err)
			return
		}

		err = h.Service.Edit.EditUserData(cookie.Value, userEdit, id)
		if err != nil {
			errorHandler(c, 400, err)
			return
		}
		c.JSON(200, gin.H{"Result": "Succes edited"})
	}
}
