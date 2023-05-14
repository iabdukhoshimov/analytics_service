package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (h *handler) setUpApi() {
	v1 := h.engine.Group("/v1")
	// v1.Use(h.auth.CheckPermissions())
	{

		v1.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	h.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
