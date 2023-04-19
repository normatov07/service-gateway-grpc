package api

import "gateway/util"

type Server struct {
	config util.Config
}

func NewServer(config util.Config) *Server {
	return &Server{
		config: config,
	}
}
