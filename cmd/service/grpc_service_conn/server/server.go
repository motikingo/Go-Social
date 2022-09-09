package server

import (
	"context"

	"github.com/motikingo/websocketproject/cmd/service"
	pb "github.com/motikingo/websocketproject/cmd/service/grpc_service_conn/proto"
	"github.com/motikingo/websocketproject/internal/pkg/entity"
	"github.com/motikingo/websocketproject/package/Helper"
)

type GrpcEEHandler struct {
	// Servers     []string
	MainService *service.MainService
}

func NewGrpcEEServer(mainService *service.MainService) *GrpcEEHandler {
	return &GrpcEEHandler{
		// Servers:     []string{"127.0.0.1:8081"},
		MainService: mainService,
	}
}

func (srv *GrpcEEHandler) GetServers() {

}
func (srv *GrpcEEHandler) HandleEEMessage(ctx context.Context, req *pb.EEBinary) (*pb.EEResponse, error) {
	println("RPC SERVER : I have handled One Request ")
	if srv.MainService != nil {
		// Creating the EEBinary
		eebin := entity.EEMBinary{
			UserID: req.UserID,
			Data:   Helper.MarshalThis(req.Data),
		}
		srv.MainService.ADDEEMBinary <- eebin
		return &pb.EEResponse{Success: true}, nil
	}
	return &pb.EEResponse{Success: false}, nil
}