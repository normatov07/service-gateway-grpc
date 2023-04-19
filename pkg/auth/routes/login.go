package routes

import (
	"context"
	"fmt"
	"gateway/pkg/auth/pb"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context, client pb.AuthServceClient) {
	req := LoginRequestBody{}

	if err := ctx.BindJSON(&req); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("Error binding request: %v", err))
		return
	}
	fmt.Println(req)
	fmt.Println(req)
	fmt.Println(req)
	res, err := client.Login(context.Background(), &pb.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, fmt.Errorf("Error while login: %v", err))
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
