package auth

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"slices"
	"strings"
)

type AuthInterceptor struct {
	jwtManager      *JWTManager
	accessibleRoles []string
}

// NewAuthInterceptor returns a new auth interceptor
func NewAuthInterceptor(jwtManager *JWTManager, accessibleRoles []string) *AuthInterceptor {
	return &AuthInterceptor{jwtManager, accessibleRoles}
}

// Unary returns a server interceptor function to authenticate and authorize unary RPC
func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		log.Println("--> unary interceptor: ", info.FullMethod)

		err := interceptor.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}

// Stream returns a server interceptor function to authenticate and authorize stream RPC
func (interceptor *AuthInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		log.Println("--> stream interceptor: ", info.FullMethod)

		err := interceptor.authorize(stream.Context(), info.FullMethod)
		if err != nil {
			return err
		}

		return handler(srv, stream)
	}
}

func (interceptor *AuthInterceptor) authorize(ctx context.Context, method string) error {

	if !slices.Contains(interceptor.accessibleRoles, method) {
		return nil
	}

	md, ok := metadata.FromIncomingContext(ctx)
	log.Println("--> unary interceptor: ", md)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := values[0]

	prefix := "Bearer "
	if !strings.HasPrefix(accessToken, prefix) {
		return status.Errorf(codes.Unauthenticated, "bearer token is not provided")
	}

	claims, err := interceptor.jwtManager.Verify(strings.TrimPrefix(accessToken, prefix))
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	if claims.Valid() != nil {
		return status.Error(codes.PermissionDenied, "token invalid")

	}

	return nil

	//if slices.Contains(claims.Service, method) {
	//	return nil
	//}
	//
	//return status.Error(codes.PermissionDenied, "no permission to access this RPC")
}
