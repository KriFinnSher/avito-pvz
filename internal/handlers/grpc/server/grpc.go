package server

import (
	grpc2 "avito-pvz/internal/handlers/grpc"
	"avito-pvz/internal/usecase"
	pb "avito-pvz/proto/pvz_v1"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

func RunGRPCServer(pvzUseCase *usecase.PVZUseCase, logger *slog.Logger) {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		logger.Error("failed to listen", "err", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPVZServiceServer(grpcServer, grpc2.NewPVZHandler(pvzUseCase))

	go func() {
		logger.Info("gRPC server starting on port 3000")
		if err = grpcServer.Serve(lis); err != nil {
			logger.Error("gRPC server error", "err", err)
		}
	}()
}
