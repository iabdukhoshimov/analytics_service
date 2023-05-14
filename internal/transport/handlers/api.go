package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (h *handler) setUpApi() {
	v1 := h.engine.Group("/v1")
	v1.Use(h.auth.CheckPermissions())
	{
		v1.POST("/login", h.login)
		v1.GET("/auth/me", h.getUserByToken)

		v1.POST("/users", h.createUser)
		v1.GET("/users/:id", h.getUserByID)
		v1.GET("/users", h.getUsers)
		v1.PUT("/users/:id", h.updateUserByID)
		v1.DELETE("/users/:id", h.deleteUserByID)

		v1.GET("/licenses", h.getAllLicenses)

		v1.GET("/organizations", h.getOrganizations)

		v1.GET("/declarations", h.getAllDeclarations)

		v1.GET("/payments", h.getPayments)

		v1.GET("/types/payments", h.getPaymentTypes)
		v1.GET("/types/licenses", h.getLicenseTypes)
		v1.GET("/types/statuses", h.getStatuses)

		v1.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
	}

	h.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
