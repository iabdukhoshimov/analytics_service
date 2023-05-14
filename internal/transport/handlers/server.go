package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"gitlab.com/greatsoft/xif-backend/api/openapi"
	"gitlab.com/greatsoft/xif-backend/internal/config"
	"gitlab.com/greatsoft/xif-backend/internal/core/repository"
	"gitlab.com/greatsoft/xif-backend/internal/core/services"
	"gitlab.com/greatsoft/xif-backend/internal/pkg/auth/middleware"
	"gitlab.com/greatsoft/xif-backend/internal/pkg/logger"
	"gitlab.com/greatsoft/xif-backend/pkg/jwt"
)

type Server interface {
	Run()
	Stop()
}

type handler struct {
	engine   *gin.Engine
	services *services.Service
	config   *config.Config

	auth *middleware.CustomAuthorizer
	jwt  jwt.Authenticator

	repo repository.Store
}

func NewServer(cfg *config.Config) Server {
	handler := handler{}

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Project.Timeout)
	defer cancel()

	handler.config = cfg

	switch cfg.Project.Mode {
	case "dev":
		gin.SetMode(gin.DebugMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	handler.engine = gin.Default()

	repos := repository.New(ctx, cfg)

	handler.repo = repos
	handler.setUpAuth()

	handler.services = services.New(repos)
	handler.setUpApi()

	return &handler
}

// @title Backend App
// @description This API contains the source for the backend app
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @BasePath /v1

// Run initializes http server
func (h *handler) Run() {
	if h.config.Project.SwaggerEnabled {
		openapi.SwaggerInfo.Version = h.config.Project.Version

		ginSwagger.WrapHandler(swaggerFiles.Handler,
			ginSwagger.URL(fmt.Sprintf(
				"%s/%d/swagger/docs.json",
				h.config.Http.Host,
				h.config.Http.Port,
			)),
			ginSwagger.DefaultModelsExpandDepth(-1),
		)
	}

	h.engine.Run(fmt.Sprintf("%s:%d", h.config.Http.Host, h.config.Http.Port))
}

func (h *handler) Stop() {
	logger.Log.Info("shutting down")
}

func (h *handler) setUpAuth() {
	tokenTypes := map[jwt.TokenType]jwt.Cred{
		jwt.AccessToken: {
			Secret: h.config.JWT.JwtSecret,
			Expire: h.config.JWT.AccessTokenTTLMinutes,
		},
		jwt.RefreshToken: {
			Secret: h.config.JWT.RefreshSecret,
			Expire: h.config.JWT.RefreshTokenTTLHours,
		},
	}

	h.jwt = jwt.New(tokenTypes)
	h.auth = middleware.New(h.jwt, h.repo)
}
