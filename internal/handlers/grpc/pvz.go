package grpc

import (
	"avito-pvz/internal/usecase"
	pb "avito-pvz/proto/pvz_v1"
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type PVZHandler struct {
	pb.UnimplementedPVZServiceServer
	pvzUseCase *usecase.PVZUseCase
}

func NewPVZHandler(pvzUseCase *usecase.PVZUseCase) *PVZHandler {
	return &PVZHandler{
		pvzUseCase: pvzUseCase,
	}
}

func (h *PVZHandler) GetPVZList(ctx context.Context, _ *pb.GetPVZListRequest) (*pb.GetPVZListResponse, error) {
	pvzs, err := h.pvzUseCase.GetAllPVZs(ctx)
	if err != nil {
		return nil, err
	}

	var result []*pb.PVZ
	for _, p := range pvzs {
		result = append(result, &pb.PVZ{
			Id:               p.ID.String(),
			RegistrationDate: timestamppb.New(p.RegistrationDate),
			City:             p.City,
		})
	}

	return &pb.GetPVZListResponse{Pvzs: result}, nil
}
