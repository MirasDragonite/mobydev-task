package handlers

import (
	"miras/internal/services"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Gin     *gin.Engine
	Service *services.Services
}

func NewHandler(service *services.Services) *Handler {
	return &Handler{Service: service, Gin: gin.Default()}
}

func (h *Handler) Router() {
	h.Gin.POST("/auth/sign-up", h.signUp())
	h.Gin.POST("/auth/sign-in", h.signIn())
	h.Gin.GET("/edit", h.getUserDataForEdit())
	h.Gin.POST("/edit", h.editUserData())
	h.Gin.POST("/logout", h.logout())
}
