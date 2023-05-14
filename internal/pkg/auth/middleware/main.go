package middleware

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"

// 	"gitlab.com/voxe-analytics/internal/pkg/logger"
// 	"gitlab.com/voxe-analytics/internal/pkg/response"
// 	"gitlab.com/voxe-analytics/pkg/jwt"
// )

// type Authorizer interface {
// 	// PermissionGetOneByRoleAndPath(ctx context.Context, arg sqlc.PermissionGetOneByRoleAndPathParams) (sqlc.PermissionGetOneByRoleAndPathRow, error)
// }

// type CustomAuthorizer struct {
// 	jwt jwt.Authenticator
// 	db  Authorizer
// }

// func New(jwt jwt.Authenticator, db Authorizer) *CustomAuthorizer {
// 	return &CustomAuthorizer{
// 		jwt: jwt,
// 		db:  db,
// 	}
// }

// func (obj *CustomAuthorizer) CheckPermissions() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// path := c.Request.URL.Path
// 		authInfo, err := obj.jwt.GetAuthInfo(c)
// 		if err != nil {
// 			logger.Log.Error("error while fetching user role from request: ", err)
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Error{Message: "Unauthorized"})
// 			return
// 		}

// 		// permission, err := obj.db.PermissionGetOneByRoleAndPath(
// 		// 	c.Request.Context(),
// 		// 	sqlc.PermissionGetOneByRoleAndPathParams{
// 		// 		RoleID: authInfo.Role,
// 		// 		Path:   path,
// 		// 	},
// 		// )

// 		// if err != nil {
// 		// 	if err != pgx.ErrNoRows {
// 		// 		logger.Log.Error("error while fetching permission from db: ", err)
// 		// 		c.AbortWithStatusJSON(http.StatusInternalServerError, response.Error{Message: "Internal Server Error"})
// 		// 		return
// 		// 	}

// 		// 	logger.Log.Warn("user does not have permission to access this resource")
// 		// 	c.AbortWithStatusJSON(http.StatusForbidden, response.Error{Message: "Forbidden"})
// 		// 	return
// 		// }

// 		// action := obj.methodToAction(c.Request.Method)
// 		// if !obj.hasPermission(&permission, action) {
// 		// 	logger.Log.Warn("user does not have permission to access this resource")
// 		// 	c.AbortWithStatusJSON(http.StatusForbidden, response.Error{Message: "Forbidden"})
// 		// 	return
// 		// }

// 		c.Set("id", authInfo.ID)
// 		c.Set("role_id", authInfo.Role)

// 		c.Next()
// 	}
// }

// func (obj *CustomAuthorizer) methodToAction(method string) string {
// 	switch method {
// 	case http.MethodGet:
// 		return "read"
// 	case http.MethodPost:
// 		return "insert"
// 	case http.MethodPut:
// 		return "update"
// 	case http.MethodDelete:
// 		return "delete"
// 	default:
// 		return ""
// 	}
// }

// // func (obj *CustomAuthorizer) hasPermission(permission *sqlc.PermissionGetOneByRoleAndPathRow, action string) bool {
// // 	switch action {
// // 	case "read":
// // 		return permission.CanRead
// // 	case "insert":
// // 		return permission.CanInsert
// // 	case "update":
// // 		return permission.CanUpdate
// // 	case "delete":
// // 		return permission.CanDelete
// // 	default:
// // 		return false
// // 	}
// // }
