package auth

import (
	"context"

	ssov1 "github.com/justcgh9/zvukovat/protos/gen/go/sso"
	"github.com/justcgh9/zvukovat/services/users/internal/lib/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
    emptyValue = 0
)

type serverAPI struct {
    ssov1.UnimplementedAuthServer
    secret string
}

func Register(gRPC *grpc.Server, secret string) {
    ssov1.RegisterAuthServer(gRPC, &serverAPI{
        secret: secret,
    })
}

func (s *serverAPI) ValidateToken(
    ctx context.Context,
    req *ssov1.ValidateTokenRequest,
) (*ssov1.ValidateTokenResponse, error) {
    if req.Token == "" {
        return nil, status.Error(codes.InvalidArgument, "token is required")
    }

    claims, err := jwt.ValidateAccessToken(req.Token, s.secret)
    if err != nil {
        return nil, status.Error(codes.Unauthenticated, err.Error())
    }

    usr := &ssov1.User{
        Id: claims.Payload.Id,
        Username: claims.Payload.Username,
        Email: claims.Payload.Email,
        IsActivated: claims.Payload.IsActivated,
        FavouriteTracks: claims.Payload.FavouriteTracks,
    }

    return &ssov1.ValidateTokenResponse{
        Payload: usr,
        Sub: claims.Subject,
        Iss: claims.Issuer,
        Aud: claims.Audience,
        Jti: claims.ID,
        Exp: timestamppb.New(claims.ExpiresAt.Time),
        Iat: nil,
        Nbf: nil,
    }, nil
}

