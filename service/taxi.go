package service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/Jamshid-Ismoilov/mytaxi-mvp/genproto"
	l "github.com/Jamshid-Ismoilov/mytaxi-mvp/pkg/logger"
	"github.com/Jamshid-Ismoilov/mytaxi-mvp/storage"
	"github.com/gofrs/uuid"
)

// TaxiService is an object that implements TaxiServiceServer interface in genproto.
type TaxiService struct {
	storage storage.IStorage
	logger  l.Logger
}

// NewTaxiService ...
func NewTaxiService(storage storage.IStorage, log l.Logger) *TaxiService {
	return &TaxiService{
		storage: storage,
		logger:  log,
	}
}

func (t *TaxiService) CreateClient(ctx context.Context, req *pb.Client) (*pb.Client, error) {
	
	id, err := uuid.NewV4()
	if err != nil {
		t.logger.Error("failed while generating uuid", l.Error(err))
		return nil, status.Error(codes.Internal, "failed generate uuid")
	}
	req.Id = id.String()

	client, err := t.storage.Taxi().CreateClient(*req)
	if err != nil {
		t.logger.Error("failed to create task", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to create task")
	}

	return &client, nil
}

func (t *TaxiService) GetClient(ctx context.Context, req *pb.ByIdReq) (*pb.Client, error) {
	client, err := t.storage.Taxi().GetClient(req.GetId())
	if err != nil {
		t.logger.Error("failed to get task", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to get task")
	}

	return &client, nil
}

func (t *TaxiService) UpdateClient(ctx context.Context, req *pb.Client) (*pb.Client, error) {
	client, err := t.storage.Taxi().UpdateClient(*req)
	if err != nil {
		t.logger.Error("failed to update task", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to update task")
	}

	return &client, nil
}

func (t *TaxiService) DeleteClient(ctx context.Context, req *pb.ByIdReq) (*pb.EmptyResp, error) {
	err := t.storage.Taxi().DeleteClient(req.Id)
	if err != nil {
		t.logger.Error("failed to delete task", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to delete task")
	}

	return &pb.EmptyResp{}, nil
}


func (t *TaxiService) CreateDriver(ctx context.Context, req *pb.Driver) (*pb.Driver, error) {
	
	id, err := uuid.NewV4()
	if err != nil {
		t.logger.Error("failed while generating uuid", l.Error(err))
		return nil, status.Error(codes.Internal, "failed generate uuid")
	}
	req.Id = id.String()

	driver, err := t.storage.Taxi().CreateDriver(*req)
	if err != nil {
		t.logger.Error("failed to create task", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to create task")
	}

	return &driver, nil
}

func (t *TaxiService) GetDriver(ctx context.Context, req *pb.ByIdReq) (*pb.Driver, error) {
	driver, err := t.storage.Taxi().GetDriver(req.GetId())
	if err != nil {
		t.logger.Error("failed to get task", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to get task")
	}

	return &driver, nil
}

func (t *TaxiService) UpdateDriver(ctx context.Context, req *pb.Driver) (*pb.Driver, error) {
	driver, err := t.storage.Taxi().UpdateDriver(*req)
	if err != nil {
		t.logger.Error("failed to update task", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to update task")
	}

	return &driver, nil
}

func (t *TaxiService) DeleteDriver(ctx context.Context, req *pb.ByIdReq) (*pb.EmptyResp, error) {
	err := t.storage.Taxi().DeleteDriver(req.Id)
	if err != nil {
		t.logger.Error("failed to delete task", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to delete task")
	}

	return &pb.EmptyResp{}, nil
}


func (t *TaxiService) CreateOrder(ctx context.Context, req *pb.Order) (*pb.FullOrder, error) {	
	id, err := uuid.NewV4()
	if err != nil {
		t.logger.Error("failed while generating uuid", l.Error(err))
		return nil, status.Error(codes.Internal, "failed generate uuid")
	}
	req.Id = id.String()

	fullorder, err := t.storage.Taxi().CreateOrder(*req)
	if err != nil {
		t.logger.Error("failed to create task", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to create task")
	}

	return &fullorder, nil
}

func (t *TaxiService) GetOrder(ctx context.Context, req *pb.ByIdReq) (*pb.FullOrder, error) {
	order, err := t.storage.Taxi().GetOrder(req.GetId())
	if err != nil {
		t.logger.Error("failed to get task", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to get task")
	}

	return &order, nil
}

func (t *TaxiService) UpdateOrder(ctx context.Context, req *pb.Order) (*pb.Order, error) {
	order, err := t.storage.Taxi().UpdateOrder(*req)
	if err != nil {
		t.logger.Error("failed to update task", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to update task")
	}

	return &order, nil
}

func (t *TaxiService) DeleteOrder(ctx context.Context, req *pb.ByIdReq) (*pb.EmptyResp, error) {
	err := t.storage.Taxi().DeleteOrder(req.Id)
	if err != nil {
		t.logger.Error("failed to delete task", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to delete task")
	}

	return &pb.EmptyResp{}, nil
}

func (t *TaxiService) ListOrder(ctx context.Context, req *pb.ListReq) (*pb.ListResp, error) {
	orders, count, err := t.storage.Taxi().ListOrder(req.ClientId, req.From, req.To, req.Page, req.Limit)
	if err != nil {
		t.logger.Error("failed to list tasks", l.Error(err))
		return nil, status.Error(codes.Internal, "failed to list tasks")
	}

	return &pb.ListResp{
		Orders: orders,
		Count: count,
	}, nil
}
