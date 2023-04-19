package routes

import (
	"context"
	"fmt"
	"gateway/pkg/auth/pb"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(ctx *gin.Context, client pb.AuthServceClient) {
	var req RegisterRequestBody

	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("Error binding request: %v", err))
		return
	}
	fmt.Println(req)
	res, err := client.Register(context.Background(), &pb.RegisterRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, fmt.Errorf("Error while register: %v", err))
		return
	}
	ctx.JSON(http.StatusOK, res)
}
