package auth

import (
	"gateway/pkg/auth/routes"
	"gateway/util"

	"github.com/gin-gonic/gin"
)

func RegisterRoute(c *util.Config, r *gin.Engine) *ServiceClient {
	src := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/auth")
	routes.POST("/register", src.Register)
	routes.POST("/login", src.Login)

	return src
}

func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}
