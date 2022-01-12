package storage

import (
	"github.com/jmoiron/sqlx"

	"github.com/Jamshid-Ismoilov/mytaxi-mvp/storage/postgres"
	"github.com/Jamshid-Ismoilov/mytaxi-mvp/storage/repo"
)

// IStorage ...
type IStorage interface {
	Taxi() repo.TaxiStorageI
}

type storagePg struct {
	db       *sqlx.DB
	taxiRepo repo.TaxiStorageI
}

// NewStoragePg ...
func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:       db,
		taxiRepo: postgres.NewTaxiRepo(db),
	}
}

func (s storagePg) Taxi() repo.TaxiStorageI {
	return s.taxiRepo
}
