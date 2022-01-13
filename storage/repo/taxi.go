package repo

import (
	pb "github.com/Jamshid-Ismoilov/mytaxi-mvp/genproto"
)

// Taxi Storage interface

type TaxiStorageI interface {
	CreateClient(pb.Client) (pb.Client, error)
	DeleteClient(id string) error
	GetClient(id string) (pb.Client, error)
	UpdateClient(pb.Client) (pb.Client, error)

	CreateDriver(pb.Driver) (pb.Driver, error)
	DeleteDriver(id string) error
	GetDriver(id string) (pb.Driver, error)
	UpdateDriver(pb.Driver) (pb.Driver, error)

	CreateOrder(pb.Order) (pb.FullOrder, error)
	DeleteOrder(id string) error
	GetOrder(id string) (pb.FullOrder, error)
	ListOrder(clientid, from, to string, page, limit int64) ([]*pb.Order, int64, error)
	UpdateOrder(pb.Order) (pb.Order, error)
}
