package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/mseptian/qbit/internal/config"
	"github.com/mseptian/qbit/internal/handler"
	"github.com/mseptian/qbit/internal/pb"
	"github.com/mseptian/qbit/internal/repository"
	"github.com/mseptian/qbit/internal/service"
	"github.com/mseptian/qbit/pkg/auth"
	"github.com/mseptian/qbit/pkg/database"
	"github.com/mseptian/qbit/pkg/hash"
	"github.com/mseptian/qbit/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen(config.DefaultGRPCServerConfig.Network, config.DefaultGRPCServerConfig.Address)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed to listen: %s", err.Error()))
	}

	//if enableTLS {
	//	tlsCredentials, err := loadTLSCredentials()
	//	if err != nil {
	//		return fmt.Errorf("cannot load TLS credentials: %w", err)
	//	}
	//
	//	serverOptions = append(serverOptions, grpc.Creds(tlsCredentials))
	//}

	db, err := database.ConnectDB()

	if err != nil {
		logger.Error("error", zap.Error(err))
	}

	hashier := hash.NewHashingPassword()
	repositories := repository.NewRepository(db)
	token := auth.NewJWTManager(os.Getenv("JWT_SECRET"), time.Hour*24)

	interceptor := auth.NewAuthInterceptor(token, []string{
		pb.User_GetUser_FullMethodName,
		pb.Cart_GetCart_FullMethodName,
		pb.Cart_PostAddToCart_FullMethodName,
		pb.Cart_GetCheckout_FullMethodName,
	})
	serverOptions := []grpc.ServerOption{
		grpc.UnaryInterceptor(interceptor.Unary()),
	}
	s := grpc.NewServer(serverOptions...)

	services := service.NewService(service.Deps{
		Repository: repositories,
		Hashing:    *hashier,
		Token:      *token,
	})

	pb.RegisterAuthServer(s, handler.NewAuthHandlerGrpcHandler(services.Auth))
	pb.RegisterUserServer(s, handler.NewUserHandlerGrpcHandler(services.User))
	pb.RegisterProductServer(s, handler.NewProductHandlerGrpcHandler(services.Product))
	pb.RegisterCartServer(s, handler.NewCartHandlerGrpcHandler(services.Cart))

	log.Println("Serving gRPC on connection ")
	go func() {
		logger.Fatal(fmt.Sprintf("%s", s.Serve(lis).Error()))
	}()

	// Create a client connection to the gRPC server we just started
	conn, err := grpc.Dial(config.DefaultGRPCServerConfig.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed to dial server: %s", err.Error()))
	}
	defer conn.Close()

	mux := runtime.NewServeMux()
	pb.RegisterAuthHandler(context.Background(), mux, conn)
	pb.RegisterUserHandler(context.Background(), mux, conn)
	pb.RegisterProductHandler(context.Background(), mux, conn)
	pb.RegisterCartHandler(context.Background(), mux, conn)

	gwServer := &http.Server{
		Addr:    config.DefaultReverseProxyConfig.Address,
		Handler: mux,
	}

	logger.Info("Serving gRPC-Gateway on connection")
	log.Fatalln(gwServer.ListenAndServe())
}
