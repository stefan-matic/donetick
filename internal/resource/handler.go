package resource

import (
	"donetick.com/core/config"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)
type Resource struct {
	Idp identityProvider `json:"identity_provider" binding:"omitempty"`
}
type identityProvider struct {
	Auth_url  string `json:"auth_url" binding:"omitempty"`
	Client_ID string `json:"client_id" binding:"omitempty"`
	Name      string `json:"name" binding:"omitempty"`
}

type Handler struct {
	config config.Config
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{
		config: *cfg,
	}
}

func (h *Handler) getResource(c *gin.Context) {
	c.JSON(200, &Resource{
		Idp: identityProvider{
		// skip resource endpoint for donetick.com
		resourceRoutes.GET("", func(c *gin.Context) {
			c.JSON(200, gin.H{})
		})
		return
	}

		resourceRoutes.GET("", h.getResource)
	}

}
