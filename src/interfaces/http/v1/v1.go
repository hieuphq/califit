package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/hieuphq/califit/src/interfaces/http/v1/user"
	"github.com/hieuphq/califit/src/registry"
)

// Apply route
func Apply(r *gin.RouterGroup, ctn registry.Container) {
	userH := user.CreateHandler(ctn.UserUC())
	r.POST("/users", userH.CreateUserHandler)
	// authH := auth.NewHandler(ctn.AuthUsecase(), ctn.JwtHelper())
	// v1 := r.Group("/v1")
	// {
	// 	auth := v1.Group("/auth")
	// 	auth.POST("/login", authH.LoginHandler)
	// 	auth.POST("/refresh-token", authH.RefreshTokenHandler)

	// 	user := v1.Group("/users")
	// 	amw := middleware.NewAuthMiddleware(ctn.JwtHelper(), ctn.RoleManager())
	// 	user.Use(amw.Auth())
	// 	{
	// 		user.GET("/me", helloHandler)
	// 	}
	// }

}

// func helloHandler(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"text": "Hello World.",
// 	})
// }
