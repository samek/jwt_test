package main

import (
	"context"
	"jwt_test/jwt"
)

type GrpcRecServer struct {
}

func (g GrpcRecServer) GenerateJwt(ctx context.Context, request *jwt.JwtGenerateRequest) (*jwt.JwtResponse, error) {
	panic("implement me")
}

func (g GrpcRecServer) CheckJwt(ctx context.Context, check *jwt.JwtCheck) (*jwt.JwtResponse, error) {
	panic("implement me")
}

func (g GrpcRecServer) GetHealth(ctx context.Context, blank *jwt.Blank) (*jwt.HealthResponse, error) {
	panic("implement me")
}


