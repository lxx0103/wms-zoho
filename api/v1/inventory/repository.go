package inventory

import (
	"strings"

	"github.com/jmoiron/sqlx"
)

type inventoryRepository struct {
	conn *sqlx.DB
}

func NewInventoryRepository(connection *sqlx.DB) InventoryRepository {
	return &inventoryRepository{
		conn: connection,
	}
}

type InventoryRepository interface {
	//Item Management
	GetItemByID(id int64) (Item, error)
	GetItemCount(filter ItemFilter) (int, error)
	GetItemList(filter ItemFilter) ([]Item, error)
}

func (r *inventoryRepository) GetItemByID(id int64) (Item, error) {
	var item Item
	err := r.conn.Get(&item, "SELECT * FROM i_items WHERE id = ? ", id)
	if err != nil {
		return Item{}, err
	}
	return item, nil
}

func (r *inventoryRepository) GetItemCount(filter ItemFilter) (int, error) {
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.SKU; v != "" {
		where, args = append(where, "sku like ?"), append(args, "%"+v+"%")
	}
	if v := filter.Name; v != "" {
		where, args = append(where, "name like ?"), append(args, "%"+v+"%")
	}
	var count int
	err := r.conn.Get(&count, `
		SELECT count(1) as count 
		FROM i_items 
		WHERE `+strings.Join(where, " AND "), args...)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *inventoryRepository) GetItemList(filter ItemFilter) ([]Item, error) {
	where, args := []string{"1 = 1"}, []interface{}{}
	if v := filter.SKU; v != "" {
		where, args = append(where, "sku like ?"), append(args, "%"+v+"%")
	}
	if v := filter.Name; v != "" {
		where, args = append(where, "name like ?"), append(args, "%"+v+"%")
	}
	args = append(args, filter.PageId*filter.PageSize-filter.PageSize)
	args = append(args, filter.PageSize)
	var items []Item
	err := r.conn.Select(&items, `
		SELECT * 
		FROM i_items 
		WHERE `+strings.Join(where, " AND ")+`
		LIMIT ?, ?
	`, args...)
	if err != nil {
		return nil, err
	}
	return items, nil
}
