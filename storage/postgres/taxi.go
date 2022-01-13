package postgres

import (
	"database/sql"
	"time"

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
	result, err := t.db.Exec(`UPDATE clients SET fullname=$2, phone=$3 WHERE id=$1`,
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
	result, err := t.db.Exec(`UPDATE drivers SET fullname=$2, phone=$3 WHERE id=$1`,
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


func (t *taxiRepo) CreateOrder(order pb.Order) (pb.FullOrder, error) {

	var id string
	err := t.db.QueryRow(`
        INSERT INTO orders(id, driver_id, client_id, status)
        VALUES ($1,$2,$3, $4) returning id`, order.Id, order.DriverId, order.ClientId, order.Status).Scan(&id)
	if err != nil {
		return pb.FullOrder{}, err
	}
	fullOrder, err := t.GetOrder(id)

	if err != nil {
		return pb.FullOrder{}, err
	}

	return fullOrder, nil
}

func (t *taxiRepo) GetOrder(id string) (pb.FullOrder, error) {

	var fullOrder pb.FullOrder
	var driverId, clientId string

	err := t.db.QueryRow(`
        SELECT id, driver_id, client_id, status FROM orders
        WHERE id=$1 and deleted_at is null`, id).Scan(
			&fullOrder.Id, 
			&driverId, 
			&clientId, 
			&fullOrder.Status, 
		)

	if err != nil {
		return pb.FullOrder{}, err
	}

	driver, err := t.GetDriver(driverId)
	if err != nil {
		return pb.FullOrder{}, err
	}

	client, err := t.GetClient(clientId)
	if err != nil {
		return pb.FullOrder{}, err
	}

	fullOrder.Driver = &driver
	fullOrder.Client = &client

	return fullOrder, nil
}


func (t *taxiRepo) UpdateOrder(order pb.Order) (pb.Order, error) {
	var status string

	err := t.db.QueryRow(`SELECT status from orders WHERE id=$1`,order.Id).Scan(&status)
	if err != nil {
		return pb.Order{}, err
	}

	if status == "accepted" && order.Status == "cancelled" || status == "finished" && order.Status == "cancelled" {
		return pb.Order{}, sql.ErrNoRows
	}  

	result, err := t.db.Exec(`UPDATE orders SET status=$2, updated_at=current_timestamp WHERE id=$1`,
		order.Id, order.Status)
	if err != nil {
		return pb.Order{}, err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return pb.Order{}, sql.ErrNoRows
	}

	res, err := t.GetOrder(order.Id)
	if err != nil {
		return pb.Order{}, err
	}
	
	order.Status = res.Status
	order.DriverId = res.Client.Id
	order.ClientId = res.Client.Id
	
	return order, nil
}

func (t *taxiRepo) DeleteOrder(id string) error {
	result, err := t.db.Exec(`UPDATE orders SET deleted_at = current_timestamp WHERE id=$1`, id)
	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (t *taxiRepo) ListOrder(clientId string, from, to string, page, limit int64) ([]*pb.Order, int64, error) {
	offset := (page - 1) * limit
	
	if from == "" || to == "" {
		from = "2000-01-01"
		to = time.Now().AddDate(0, 0, 1).Format("2006-01-02")
	}
	
	rows, err := t.db.Queryx(
		`SELECT id, driver_id, client_id, status FROM orders WHERE client_id = $1 and created_at >= $2 and created_at <= $3 LIMIT $4 OFFSET $5`,
		clientId, from, to, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	if err = rows.Err(); err != nil {
		return nil, 0, err
	}
	defer rows.Close() // nolint:errcheck

	var (
		orders []*pb.Order
		order  pb.Order
		count int64
	)
	for rows.Next() {
		err = rows.Scan(&order.Id, &order.DriverId, &order.ClientId, &order.Status)
		if err != nil {
			return nil, 0, err
		}
		orders = append(orders, &order)
	}

	err = t.db.QueryRow(
		`SELECT count(*) FROM orders WHERE client_id = $1 and created_at >= $2 and created_at <= $3`,
		clientId, from, to).Scan(&count)
	if err != nil {
		return nil, 0, err
	}
	return orders, count, nil
}
