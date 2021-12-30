package setting

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
)

type settingRepository struct {
	conn *sqlx.DB
}

func NewSettingRepository(connection *sqlx.DB) SettingRepository {
	return &settingRepository{
		conn: connection,
	}
}

type SettingRepository interface {
	//Shelf Management
	GetShelfByID(id int64) (Shelf, error)
	CreateShelf(info ShelfNew) (int64, error)
	GetShelfCount(filter ShelfFilter) (int, error)
	GetShelfList(filter ShelfFilter) ([]Shelf, error)
	UpdateShelf(id int64, info ShelfNew) (int64, error)

	//Location Management
	GetLocationByID(id int64) (Location, error)
	GetLocationByCode(code string) (Location, error)
	CreateLocation(info LocationNew) (int64, error)
	GetLocationCount(filter LocationFilter) (int, error)
	GetLocationList(filter LocationFilter) ([]Location, error)
	UpdateLocation(id int64, info LocationNew) (int64, error)
	GetLocationBySKU(sku string) (*[]Location, error)
	UpdateLocationStock(UpdateLocationStock) (int64, error)
	StockTransfer(LocationStockTransfer, string) (int64, error)
	GetTransferCount(filter TransferFilter) (int, error)
	GetTransferList(filter TransferFilter) ([]TransferTransaction, error)

	//Barcode Management
	GetBarcodeByID(id int64) (Barcode, error)
	CreateBarcode(info BarcodeNew) (int64, error)
	GetBarcodeCount(filter BarcodeFilter) (int, error)
	GetBarcodeList(filter BarcodeFilter) ([]Barcode, error)
	UpdateBarcode(id int64, info BarcodeNew) (int64, error)
	GetBarcodeByCode(string) (*Barcode, error)
}

func (r *settingRepository) GetShelfByID(id int64) (Shelf, error) {
	var shelf Shelf
	err := r.conn.Get(&shelf, "SELECT * FROM s_shelves WHERE id = ? ", id)
	if err != nil {
		return Shelf{}, err
	}
	return shelf, nil
}
func (r *settingRepository) CreateShelf(info ShelfNew) (int64, error) {
	tx, err := r.conn.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	result, err := tx.Exec(`
		INSERT INTO s_shelves
		(
			code,
			level,
			location,
			enabled,
			created,
			created_by,
			updated,
			updated_by
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`, info.Code, info.Level, info.Location, info.Enabled, time.Now(), info.User, time.Now(), info.User)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	tx.Commit()
	return id, nil
}

func (r *settingRepository) GetShelfCount(filter ShelfFilter) (int, error) {
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.Code; v != "" {
		where, args = append(where, "code = ?"), append(args, v)
	}
	if v := filter.Location; v != "" {
		where, args = append(where, "location = ?"), append(args, v)
	}
	var count int
	err := r.conn.Get(&count, `
		SELECT count(1) as count 
		FROM s_shelves 
		WHERE `+strings.Join(where, " AND "), args...)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *settingRepository) GetShelfList(filter ShelfFilter) ([]Shelf, error) {
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.Code; v != "" {
		where, args = append(where, "code = ?"), append(args, v)
	}
	if v := filter.Location; v != "" {
		where, args = append(where, "location = ?"), append(args, v)
	}
	args = append(args, filter.PageId*filter.PageSize-filter.PageSize)
	args = append(args, filter.PageSize)
	var shelves []Shelf
	err := r.conn.Select(&shelves, `
		SELECT * 
		FROM s_shelves 
		WHERE `+strings.Join(where, " AND ")+`
		LIMIT ?, ?
	`, args...)
	if err != nil {
		return nil, err
	}
	return shelves, nil
}

func (r *settingRepository) UpdateShelf(id int64, info ShelfNew) (int64, error) {
	tx, err := r.conn.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	result, err := tx.Exec(`
		Update s_shelves SET 
		code = ?,
		level = ?,
		location = ?,
		enabled = ?,
		updated = ?,
		updated_by = ? 
		WHERE id = ?
	`, info.Code, info.Level, info.Location, info.Enabled, time.Now(), info.User, id)
	if err != nil {
		return 0, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	tx.Commit()
	return affected, nil
}

func (r *settingRepository) GetLocationByID(id int64) (Location, error) {
	var location Location
	err := r.conn.Get(&location, "SELECT * FROM s_locations WHERE id = ? ", id)
	if err != nil {
		return Location{}, err
	}
	return location, nil
}
func (r *settingRepository) GetLocationByCode(code string) (Location, error) {
	var location Location
	err := r.conn.Get(&location, "SELECT * FROM s_locations WHERE code = ? ", code)
	if err != nil {
		return Location{}, err
	}
	return location, nil
}
func (r *settingRepository) CreateLocation(info LocationNew) (int64, error) {
	tx, err := r.conn.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	result, err := tx.Exec(`
		INSERT INTO s_locations
		(
			code,
			level,
			shelf_id,
			sku,
			capacity,
			quantity,
			available,
			alert,
			can_pick,
			unit,
			enabled,
			created,
			created_by,
			updated,
			updated_by
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, info.Code, info.Level, info.ShelfID, info.SKU, info.Capacity, info.Quantity, info.Capacity-info.Quantity, info.Alert, info.Quantity, info.Unit, info.Enabled, time.Now(), info.User, time.Now(), info.User)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	tx.Commit()
	return id, nil
}

func (r *settingRepository) GetLocationCount(filter LocationFilter) (int, error) {
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.Code; v != "" {
		where, args = append(where, "code like ?"), append(args, "%"+v+"%")
	}
	if v := filter.SKU; v != "" {
		where, args = append(where, "sku = ?"), append(args, v)
	}
	if v := filter.Level; v != "" {
		where, args = append(where, "level = ?"), append(args, v)
	}
	if v := filter.ShelfID; v != 0 {
		where, args = append(where, "shelf_id = ?"), append(args, v)
	}
	if v := filter.IsAlert; v {
		where = append(where, "quantity < alert")
	}
	if v := filter.IsActive; v {
		where = append(where, "enabled = 1")
	}
	var count int
	err := r.conn.Get(&count, `
		SELECT count(1) as count 
		FROM s_locations 
		WHERE `+strings.Join(where, " AND "), args...)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *settingRepository) GetLocationList(filter LocationFilter) ([]Location, error) {
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.Code; v != "" {
		where, args = append(where, "code like ?"), append(args, "%"+v+"%")
	}
	if v := filter.SKU; v != "" {
		where, args = append(where, "sku = ?"), append(args, v)
	}
	if v := filter.Level; v != "" {
		where, args = append(where, "level = ?"), append(args, v)
	}
	if v := filter.ShelfID; v != 0 {
		where, args = append(where, "shelf_id = ?"), append(args, v)
	}
	if v := filter.IsAlert; v {
		where = append(where, "quantity < alert")
	}
	if v := filter.IsActive; v {
		where = append(where, "enabled = 1")
	}
	args = append(args, filter.PageId*filter.PageSize-filter.PageSize)
	args = append(args, filter.PageSize)
	var locations []Location
	err := r.conn.Select(&locations, `
		SELECT * 
		FROM s_locations 
		WHERE `+strings.Join(where, " AND ")+`
		LIMIT ?, ?
	`, args...)
	if err != nil {
		return nil, err
	}
	return locations, nil
}

func (r *settingRepository) GetLocationBySKU(sku string) (*[]Location, error) {
	where, args := []string{"1 = 1"}, []interface{}{}
	where, args = append(where, "sku = ?"), append(args, sku)

	var locations []Location
	err := r.conn.Select(&locations, `
		SELECT * 
		FROM s_locations 
		WHERE `+strings.Join(where, " AND ")+`
		AND available > 0
		AND enabled = 1
		ORDER BY quantity DESC, available ASC
	`, args...)
	if err != nil {
		return nil, err
	}
	return &locations, nil
}

func (r *settingRepository) UpdateLocation(id int64, info LocationNew) (int64, error) {
	tx, err := r.conn.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	result, err := tx.Exec(`
		Update s_locations SET 
		code = ?,
		level = ?,
		shelf_id = ?,
		sku = ?,
		capacity = ?,
		quantity = ?,
		available = ?, 
		alert = ?,
		unit = ?,
		enabled = ?,
		updated = ?,
		updated_by = ? 
		WHERE id = ?
	`, info.Code, info.Level, info.ShelfID, info.SKU, info.Capacity, info.Quantity, info.Capacity-info.Quantity, info.Alert, info.Unit, info.Enabled, time.Now(), info.User, id)
	if err != nil {
		return 0, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	tx.Commit()
	return affected, nil
}

func (r *settingRepository) UpdateLocationStock(info UpdateLocationStock) (int64, error) {
	tx, err := r.conn.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	result, err := tx.Exec(`
		Update s_locations SET 
		quantity = quantity + ?,
		can_pick = can_pick + ?,
		available = available - ?, 
		updated = ?,
		updated_by = ? 
		WHERE code = ?
	`, info.Quantity, info.Quantity, info.Quantity, time.Now(), info.User, info.Code)
	if err != nil {
		return 0, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	tx.Commit()
	return affected, nil
}

func (r *settingRepository) GetBarcodeByID(id int64) (Barcode, error) {
	var barcode Barcode
	err := r.conn.Get(&barcode, "SELECT * FROM s_barcodes WHERE id = ? ", id)
	if err != nil {
		return Barcode{}, err
	}
	return barcode, nil
}
func (r *settingRepository) CreateBarcode(info BarcodeNew) (int64, error) {
	tx, err := r.conn.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	result, err := tx.Exec(`
		INSERT INTO s_barcodes
		(
			code,
			sku,
			unit,
			quantity,
			enabled,
			created,
			created_by,
			updated,
			updated_by
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, info.Code, info.SKU, info.Unit, info.Quantity, info.Enabled, time.Now(), info.User, time.Now(), info.User)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	tx.Commit()
	return id, nil
}

func (r *settingRepository) GetBarcodeCount(filter BarcodeFilter) (int, error) {
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.Code; v != "" {
		where, args = append(where, "code = ?"), append(args, v)
	}
	if v := filter.SKU; v != "" {
		where, args = append(where, "sku = ?"), append(args, v)
	}
	var count int
	err := r.conn.Get(&count, `
		SELECT count(1) as count 
		FROM s_barcodes 
		WHERE `+strings.Join(where, " AND "), args...)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *settingRepository) GetBarcodeList(filter BarcodeFilter) ([]Barcode, error) {
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.Code; v != "" {
		where, args = append(where, "code = ?"), append(args, v)
	}
	if v := filter.SKU; v != "" {
		where, args = append(where, "sku = ?"), append(args, v)
	}
	args = append(args, filter.PageId*filter.PageSize-filter.PageSize)
	args = append(args, filter.PageSize)
	var barcodes []Barcode
	err := r.conn.Select(&barcodes, `
		SELECT * 
		FROM s_barcodes 
		WHERE `+strings.Join(where, " AND ")+`
		LIMIT ?, ?
	`, args...)
	if err != nil {
		return nil, err
	}
	return barcodes, nil
}

func (r *settingRepository) UpdateBarcode(barcodeID int64, info BarcodeNew) (int64, error) {
	tx, err := r.conn.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	result, err := tx.Exec(`
		Update s_barcodes SET 
		code = ?,
		sku = ?,
		unit = ?,
		quantity = ?,
		enabled = ?,
		updated = ?,
		updated_by = ? 
		WHERE id = ?
	`, info.Code, info.SKU, info.Unit, info.Quantity, info.Enabled, time.Now(), info.User, barcodeID)
	if err != nil {
		return 0, err
	}
	id, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	tx.Commit()
	return id, nil
}

func (r *settingRepository) GetBarcodeByCode(code string) (*Barcode, error) {
	var barcode Barcode
	err := r.conn.Get(&barcode, "SELECT * FROM s_barcodes WHERE code = ? ", code)
	if err != nil {
		return nil, err
	}
	return &barcode, nil
}

func (r *settingRepository) StockTransfer(info LocationStockTransfer, user string) (int64, error) {
	tx, err := r.conn.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	var fromLocation Location
	err = r.conn.Get(&fromLocation, `SELECT * FROM s_locations WHERE code = ? LIMIT 1`, info.From)
	if err != nil {
		fmt.Println("llllll")
		return 0, err
	}
	var toLocation Location
	err = r.conn.Get(&toLocation, `SELECT * FROM s_locations WHERE code = ? LIMIT 1`, info.To)
	if err != nil {
		return 0, err
	}
	if toLocation.SKU != fromLocation.SKU {
		return 0, errors.New("LOCATION SKU NOT THE SAME")
	}
	if toLocation.Available < info.Quantity {
		return 0, errors.New("TO LOCATION NOT ENOUGH")
	}
	if fromLocation.Quantity < info.Quantity {
		return 0, errors.New("FROM LOCATION NOT ENOUGH")
	}
	result, err := tx.Exec(`
		INSERT INTO i_transfer_transactions
		(
			from_code,
			to_code,
			sku,
			quantity,
			enabled,
			created,
			created_by,
			updated,
			updated_by
		)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, info.From, info.To, fromLocation.SKU, info.Quantity, 1, time.Now(), user, time.Now(), user)
	if err != nil {
		return 0, err
	}
	transactionID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	_, err = tx.Exec(`
		Update s_locations SET 
		quantity = quantity + ?,
		available = available - ?,
		can_pick = can_pick + ?,
		updated = ?,
		updated_by = ? 
		WHERE code = ?
	`, info.Quantity, info.Quantity, info.Quantity, time.Now(), user, info.To)
	if err != nil {
		return 0, err
	}
	_, err = tx.Exec(`
		Update s_locations SET 
		quantity = quantity - ?,
		available = available + ?,
		can_pick = can_pick - ?,
		updated = ?,
		updated_by = ? 
		WHERE code = ?
	`, info.Quantity, info.Quantity, info.Quantity, time.Now(), user, info.From)
	if err != nil {
		return 0, err
	}
	tx.Commit()
	return transactionID, nil
}

func (r *settingRepository) GetTransferCount(filter TransferFilter) (int, error) {
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.From; v != "" {
		where, args = append(where, "from_code like ?"), append(args, "%"+v+"%")
	}
	if v := filter.To; v != "" {
		where, args = append(where, "to_code like ?"), append(args, "%"+v+"%")
	}
	var count int
	err := r.conn.Get(&count, `
		SELECT count(1) as count 
		FROM i_transfer_transactions 
		WHERE `+strings.Join(where, " AND "), args...)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *settingRepository) GetTransferList(filter TransferFilter) ([]TransferTransaction, error) {
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.From; v != "" {
		where, args = append(where, "from_code like ?"), append(args, "%"+v+"%")
	}
	if v := filter.To; v != "" {
		where, args = append(where, "to_code like ?"), append(args, "%"+v+"%")
	}
	args = append(args, filter.PageId*filter.PageSize-filter.PageSize)
	args = append(args, filter.PageSize)
	var transfers []TransferTransaction
	err := r.conn.Select(&transfers, `
		SELECT * 
		FROM i_transfer_transactions 
		WHERE `+strings.Join(where, " AND ")+`
		LIMIT ?, ?
	`, args...)
	if err != nil {
		return nil, err
	}
	return transfers, nil
}
