package service

import (
	"context"
	"time"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/Jamshid-Ismoilov/mytaxi-mvp/genproto"
	l "github.com/Jamshid-Ismoilov/mytaxi-mvp/pkg/logger"
	"github.com/Jamshid-Ismoilov/mytaxi-mvp/storage"
	"github.com/gofrs/uuid"
)

// TaxiService is an object that implements TaxiServiceServer interface in genproto.
type TaxiService struct {
	storage storage.TaxiStorageI
	logger  l.Logger
}

// NewTaxiService ...
func NewTaxiService(storage storage.TaxiStorageI, log l.Logger) *TaskService {
	return &TaskService{
		storage: storage,
		logger:  log,
	}
}

func (t *TaxiService) CreateClient(ctx context.Context, req *pb.Client) (*pb.Client, error) {
	
	id, err := uuid.NewV4()
	if err != nil {
		s.logger.Error("failed while generating uuid", l.Error(err))
		return nil, status.Error(codes.Internal, "failed generate uuid")
	}
	req.Id = id.String()

	client, err := s.storage.Taxi().CreateClient(*req)
	if err != nil {
		s.logger.Error("failed to create task", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to create task")
	}

	return &client, nil
}

func (t *TaxiService) GetClient(ctx context.Context, req *pb.ByIdReq) (*pb.Client, error) {
	client, err := s.storage.Taxi().GetClient(req.GetId())
	if err != nil {
		s.logger.Error("failed to get task", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to get task")
	}

	return &client, nil
}

func (t *TaxiService) UpdateClient(ctx context.Context, req *pb.Client) (*pb.Client, error) {
	client, err := s.storage.Taxi().UpdateClient(*req)
	if err != nil {
		s.logger.Error("failed to update task", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to update task")
	}

	return &client, nil
}

func (t *TaxiService) DeleteClient(ctx context.Context, req *pb.ByIdReq) (*pb.EmptyResp, error) {
	err := s.storage.Taxi().DeleteClient(req.Id)
	if err != nil {
		s.logger.Error("failed to delete task", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to delete task")
	}

	return &pb.EmptyResp{}, nil
}


func (t *TaxiService) CreateDriver(ctx context.Context, req *pb.Driver) (*pb.Driver, error) {
	
	id, err := uuid.NewV4()
	if err != nil {
		s.logger.Error("failed while generating uuid", l.Error(err))
		return nil, status.Error(codes.Internal, "failed generate uuid")
	}
	req.Id = id.String()

	driver, err := s.storage.Taxi().CreateDriver(*req)
	if err != nil {
		s.logger.Error("failed to create task", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to create task")
	}

	return &driver, nil
}

func (t *TaxiService) GetDriver(ctx context.Context, req *pb.ByIdReq) (*pb.Driver, error) {
	driver, err := s.storage.Taxi().GetDriver(req.GetId())
	if err != nil {
		s.logger.Error("failed to get task", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to get task")
	}

	return &driver, nil
}

func (t *TaxiService) UpdateDriver(ctx context.Context, req *pb.Driver) (*pb.Driver, error) {
	driver, err := s.storage.Taxi().UpdateDriver(*req)
	if err != nil {
		s.logger.Error("failed to update task", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to update task")
	}

	return &driver, nil
}

func (t *TaxiService) DeleteDriver(ctx context.Context, req *pb.ByIdReq) (*pb.EmptyResp, error) {
	err := s.storage.Taxi().DeleteDriver(req.Id)
	if err != nil {
		s.logger.Error("failed to delete task", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to delete task")
	}

	return &pb.EmptyResp{}, nil
}