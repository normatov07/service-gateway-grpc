package auth

import (
	"fmt"
	"gateway/pkg/auth/pb"
	"gateway/util"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.AuthServceClient
}

func InitServiceClient(c *util.Config) pb.AuthServceClient {
	cc, err := grpc.Dial(
		c.AUTH_SRV_ADDRESS,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		fmt.Errorf("Auth grpc Dial error: %v", err)
	}

	return pb.NewAuthServceClient(cc)
}
