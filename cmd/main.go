package main

import (
	"fmt"
	"gateway/api"
	"gateway/util"

	"gateway/pkg/auth"

	"github.com/gin-gonic/gin"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		fmt.Errorf("Failed to load config: %v", err)
	}
	server := api.NewServer(config)
	fmt.Println(server)

	router := gin.Default()

	authRes := auth.RegisterRoute(&config, router)
	fmt.Println(authRes)

	router.Run(config.SERVER_ADDRESS)
}
