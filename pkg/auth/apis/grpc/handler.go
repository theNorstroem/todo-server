package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/spf13/viper"
	"github.com/theNorstroem/todo-server/pkg/auth/core"
	"github.com/theNorstroem/todo-server/pkg/auth/core/login"
	"github.com/theNorstroem/todo-specs/dist/pb/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
)

func Handler(s *grpc.Server, domainServices core.DomainServices) {
	authpb.RegisterAuthSessionServer(s, GetServiceServer(domainServices))

}

// Register your domainServices here
func GetServiceServer(domainServices core.DomainServices) authpb.AuthSessionServer {
	var service = &AuthService{}
	service.login = domainServices.Login

	//service.logout = domainServices.Logout
	return service
}

type AuthService struct {
	authpb.UnsafeAuthSessionServer
	login login.Service
}

// Override Funktion um nicht über die default auth-middleware
func (s AuthService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, nil
}

// login
func (a AuthService) CreateAuthSession(ctx context.Context, request *authpb.CreateAuthSessionRequest) (*empty.Empty, error) {
	token, err := a.login.Login(*request.Body)
	if err == nil {
		err = grpc.SendHeader(ctx, metadata.Pairs(
			"Set-Cookie", "Authorization=Bearer "+token+"; expires="+time.Now().Add(time.Second*viper.GetDuration("server.jwt.expiry_duration_in_s")).UTC().String(),
		))
	}
	return &empty.Empty{}, err
}

//logout
func (a AuthService) DeleteAuthSession(ctx context.Context, request *authpb.DeleteAuthSessionRequest) (*empty.Empty, error) {
	err := grpc.SendHeader(ctx, metadata.Pairs("Set-Cookie", "Authorization=deleted; expires=Thu, 01 Jan 1970 00:00:00 GMT"))
	return &empty.Empty{}, err
}
