package postgres

import (
	"database/sql"
	// "time"
	// "fmt"

	"github.com/jmoiron/sqlx"

	pb "github.com/Jamshid-Ismoilov/mytaxi-mvp/genproto"
)

type taxiRepo struct {
	db *sqlx.DB
}

// NewTaxiRepo ...
func NewTaxiRepo(db *sqlx.DB) *taxiRepo {
	return &taxiRepo{db: db}
}

func (t *taxiRepo) CreateClient(client pb.Client) (pb.Client, error) {

	var id string
	err := t.db.QueryRow(`
        INSERT INTO clients(id, fullname, phone)
        VALUES ($1,$2,$3) returning id`, client.Id, client.Fullname, client.Phone).Scan(&id)
	if err != nil {
		return pb.Client{}, err
	}
	client, err = t.GetClient(id)

	if err != nil {
		return pb.Client{}, err
	}

	return client, nil
}

func (t *taxiRepo) GetClient(id string) (pb.Client, error) {

	var client pb.Client

	err := t.db.QueryRow(`
        SELECT id, fullname, phone FROM clients
        WHERE id=$1 and deleted_at is null`, id).Scan(
			&client.Id, 
			&client.Fullname, 
			&client.Phone, 
		)

	if err != nil {
		return pb.Client{}, err
	}
	return client, nil
}


func (t *taxiRepo) UpdateClient(client pb.Client) (pb.Client, error) {
	result, err := t.db.Exec(`UPDATE clients SET fullname=$2, phone=$3, updated_at=current_timestamp WHERE id=$1`,
		client.Id, client.Fullname, client.Phone)
	if err != nil {
		return pb.Client{}, err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return pb.Client{}, sql.ErrNoRows
	}

	res, err := t.GetClient(client.Id)
	if err != nil {
		return pb.Client{}, err
	}

	return res, nil
}

func (t *taxiRepo) DeleteClient(id string) error {
	result, err := t.db.Exec(`UPDATE clients SET deleted_at = current_timestamp WHERE id=$1`, id)
	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (t *taxiRepo) CreateDriver(driver pb.Driver) (pb.Driver, error) {

	var id string
	err := t.db.QueryRow(`
        INSERT INTO drivers(id, fullname, phone)
        VALUES ($1,$2,$3) returning id`, driver.Id, driver.Fullname, driver.Phone).Scan(&id)
	if err != nil {
		return pb.Driver{}, err
	}
	driver, err = t.GetDriver(id)

	if err != nil {
		return pb.Driver{}, err
	}

	return driver, nil
}

func (t *taxiRepo) GetDriver(id string) (pb.Driver, error) {

	var driver pb.Driver

	err := t.db.QueryRow(`
        SELECT id, fullname, phone FROM drivers
        WHERE id=$1 and deleted_at is null`, id).Scan(
			&driver.Id, 
			&driver.Fullname, 
			&driver.Phone, 
		)

	if err != nil {
		return pb.Driver{}, err
	}
	return driver, nil
}


func (t *taxiRepo) UpdateDriver(driver pb.Driver) (pb.Driver, error) {
	result, err := t.db.Exec(`UPDATE drivers SET fullname=$2, phone=$3, updated_at=current_timestamp WHERE id=$1`,
		driver.Id, driver.Fullname, driver.Phone)
	if err != nil {
		return pb.Driver{}, err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return pb.Driver{}, sql.ErrNoRows
	}

	driver, err = t.GetDriver(driver.Id)
	if err != nil {
		return pb.Driver{}, err
	}

	return driver, nil
}

func (t *taxiRepo) DeleteDriver(id string) error {
	result, err := t.db.Exec(`UPDATE drivers SET deleted_at = current_timestamp WHERE id=$1`, id)
	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}

