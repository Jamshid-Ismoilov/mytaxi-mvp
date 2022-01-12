package repo

import (
	pb "github.com/Jamshid-Ismoilov/mytaxi-mvp/genproto"
)

// Taxi Storage interface

type TaxiStorageI interface {
	CreateClient(pb.Client) (pb.Client, error)
	DeleteClient(pb.ByIdReq) (error)
	GetClient(pb.ByIdReq) (pb.Client, error)
	UpdateClient(pb.Client) (pb.Client, error)

	CreateDriver(pb.Driver) (pb.Driver, error)
	DeleteDriver(pb.ByIdReq) (error)
	GetDriver(pb.ByIdReq) (pb.Driver, error)
	UpdateDriver(pb.Driver) (pb.Driver, error)
	
	CreateOrder(pb.Order) (pb.Order, error)
	DeleteOrder(pb.ByIdReq) (error)
	GetOrder(pb.ByIdReq) (pb.FullOrder, error)
	ListOrder(clientId, from, to string, page limit int64) ([]*pb.Order, int64, error)
	UpdateOrder(pb.Order) (pb.Order, error)
}
