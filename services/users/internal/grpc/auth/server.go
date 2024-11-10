package auth

import (
	"context"

	ssov1 "github.com/justcgh9/zvukovat/protos/gen/go/sso"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
    emptyValue = 0
)

type serverAPI struct {
    ssov1.UnimplementedAuthServer
}

func Register(gRPC *grpc.Server) {
    ssov1.RegisterAuthServer(gRPC, &serverAPI{})
}

func (s *serverAPI) ValidateToken(
    ctx context.Context,
    req *ssov1.ValidateTokenRequest,
) (*ssov1.ValidateTokenResponse, error) {
    if req.Token == "" {
        return nil, status.Error(codes.InvalidArgument, "token is required")
    }

    return nil, nil
}

