package setting

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/tealeg/xlsx/v3"
	"wms.com/core/database"
)

type settingService struct {
}

func NewSettingService() SettingService {
	return &settingService{}
}

// SettingService represents a service for managing settings.
type SettingService interface {
	//Shelf Management
	GetShelfByID(int64) (Shelf, error)
	NewShelf(ShelfNew) (Shelf, error)
	GetShelfList(ShelfFilter) (int, []Shelf, error)
	UpdateShelf(int64, ShelfNew) (Shelf, error)
	//Location Management
	GetLocationByID(int64) (Location, error)
	NewLocation(LocationNew) (Location, error)
	GetLocationList(LocationFilter) (int, []Location, error)
	UpdateLocation(int64, LocationNew) (Location, error)
	GetLocationBySKU(string) (*[]Location, error)
	GetLocationByCode(string) (Location, error)
	UpdateLocationStock(UpdateLocationStock) (int64, error)
	StockTransfer(LocationStockTransfer, string) (int64, error)
	GetTransferList(TransferFilter) (int, []TransferTransaction, error)
	GetNextTransactionLocation(TranferFromFilter) (string, error)
	//Barcode Management
	GetBarcodeByID(int64) (Barcode, error)
	NewBarcode(BarcodeNew) (Barcode, error)
	GetBarcodeList(BarcodeFilter) (int, []Barcode, error)
	UpdateBarcode(int64, BarcodeNew) (Barcode, error)
	GetBarcodeByCode(string) (*Barcode, error)
	BatchUpload(string, string) error
	ExportBarcode(BarcodeFilterNoPage, string) error
}

func (s *settingService) GetShelfByID(id int64) (Shelf, error) {
	db := database.InitMySQL()
	repo := NewSettingRepository(db)
	shelf, err := repo.GetShelfByID(id)
	return shelf, err
}

func (s *settingService) NewShelf(info ShelfNew) (Shelf, error) {
	db := database.InitMySQL()
	repo := NewSettingRepository(db)
	shelfID, err := repo.CreateShelf(info)
	if err != nil {
		return Shelf{}, err
	}
	shelf, err := repo.GetShelfByID(shelfID)
	return shelf, err
}

func (s *settingService) GetShelfList(filter ShelfFilter) (int, []Shelf, error) {
	db := database.InitMySQL()
	repo := NewSettingRepository(db)
	count, err := repo.GetShelfCount(filter)
	if err != nil {
		return 0, nil, err
	}
	list, err := repo.GetShelfList(filter)
	if err != nil {
		return 0, nil, err
	}
	return count, list, err
}

func (s *settingService) UpdateShelf(shelfID int64, info ShelfNew) (Shelf, error) {
	db := database.InitMySQL()
	repo := NewSettingRepository(db)
	_, err := repo.UpdateShelf(shelfID, info)
	if err != nil {
		return Shelf{}, err
	}
	shelf, err := repo.GetShelfByID(shelfID)
	return shelf, err
}

func (s *settingService) GetLocationByID(id int64) (Location, error) {
	db := database.InitMySQL()
	repo := NewSettingRepository(db)
	location, err := repo.GetLocationByID(id)
	return location, err
}

func (s *settingService) NewLocation(info LocationNew) (Location, error) {
	db := database.InitMySQL()
	repo := NewSettingRepository(db)
	locationID, err := repo.CreateLocation(info)
	if err != nil {
		return Location{}, err
	}
	location, err := repo.GetLocationByID(locationID)
	return location, err
}

func (s *settingService) GetLocationList(filter LocationFilter) (int, []Location, error) {
	db := database.InitMySQL()
	repo := NewSettingRepository(db)
	count, err := repo.GetLocationCount(filter)
	if err != nil {
		return 0, nil, err
	}
	list, err := repo.GetLocationList(filter)
	if err != nil {
		return 0, nil, err
	}
	return count, list, err
}

func (s *settingService) UpdateLocation(locationID int64, info LocationNew) (Location, error) {
	db := database.InitMySQL()
	repo := NewSettingRepository(db)
	_, err := repo.UpdateLocation(locationID, info)
	if err != nil {
		return Location{}, err
	}
	location, err := repo.GetLocationByID(locationID)
	return location, err
}

func (s *settingService) GetLocationBySKU(sku string) (*[]Location, error) {
	db := database.InitMySQL()
	repo := NewSettingRepository(db)
	location, err := repo.GetLocationBySKU(sku)
	return location, err
}

func (s *settingService) GetLocationByCode(code string) (Location, error) {
	db := database.InitMySQL()
	repo := NewSettingRepository(db)
	location, err := repo.GetLocationByCode(code)
	return location, err
}
func (s settingService) UpdateLocationStock(updateInfo UpdateLocationStock) (int64, error) {
	db := database.InitMySQL()
	repo := NewSettingRepository(db)
	affected, err := repo.UpdateLocationStock(updateInfo)
	return affected, err
}

func (s *settingService) GetBarcodeByID(id int64) (Barcode, error) {
	db := database.InitMySQL()
	repo := NewSettingRepository(db)
	barcode, err := repo.GetBarcodeByID(id)
	return barcode, err
}

func (s *settingService) NewBarcode(info BarcodeNew) (Barcode, error) {
	db := database.InitMySQL()
	repo := NewSettingRepository(db)
	barcodeID, err := repo.CreateBarcode(info)
	if err != nil {
		return Barcode{}, err
	}
	barcode, err := repo.GetBarcodeByID(barcodeID)
	return barcode, err
}

func (s *settingService) GetBarcodeList(filter BarcodeFilter) (int, []Barcode, error) {
	db := database.InitMySQL()
	repo := NewSettingRepository(db)
	count, err := repo.GetBarcodeCount(filter)
	if err != nil {
		return 0, nil, err
	}
	list, err := repo.GetBarcodeList(filter)
	if err != nil {
		return 0, nil, err
	}
	return count, list, err
}

func (s *settingService) UpdateBarcode(barcodeID int64, info BarcodeNew) (Barcode, error) {
	db := database.InitMySQL()
	repo := NewSettingRepository(db)
	_, err := repo.UpdateBarcode(barcodeID, info)
	if err != nil {
		return Barcode{}, err
	}
	barcode, err := repo.GetBarcodeByID(barcodeID)
	return barcode, err
}

func (s *settingService) GetBarcodeByCode(code string) (*Barcode, error) {
	db := database.InitMySQL()
	repo := NewSettingRepository(db)
	barcode, err := repo.GetBarcodeByCode(code)
	return barcode, err
}

func (s *settingService) StockTransfer(info LocationStockTransfer, user string) (int64, error) {
	db := database.InitMySQL()
	repo := NewSettingRepository(db)
	transaction, err := repo.StockTransfer(info, user)
	return transaction, err
}

func (s *settingService) GetTransferList(filter TransferFilter) (int, []TransferTransaction, error) {
	db := database.InitMySQL()
	repo := NewSettingRepository(db)
	count, err := repo.GetTransferCount(filter)
	if err != nil {
		return 0, nil, err
	}
	list, err := repo.GetTransferList(filter)
	if err != nil {
		return 0, nil, err
	}
	return count, list, err
}

func (s *settingService) GetNextTransactionLocation(filter TranferFromFilter) (string, error) {
	db := database.InitMySQL()
	repo := NewSettingRepository(db)
	res, err := repo.GetNextTransactionLocation(filter)
	if err != nil {
		return "", err
	}
	return res, err
}

func (s *settingService) BatchUpload(path string, user string) error {
	wb, err := xlsx.OpenFile(path)
	if err != nil {
		msg := "Open excel Error!"
		return errors.New(err.Error() + msg)
	}
	sheetName := "barcodes"
	sheet, ok := wb.Sheet[sheetName]
	if !ok {
		msg := "Open Sheet Error!Sheet name must be barcodes"
		return errors.New(err.Error() + msg)
	}
	var barcodes []BarcodeNew
	sheet.ForEachRow(func(r *xlsx.Row) error {
		if r.GetCoordinate() == 0 {
			return nil
		}
		var barcode BarcodeNew
		barcode.User = user
		r.ForEachCell(func(c *xlsx.Cell) error {
			cn, rn := c.GetCoordinates()
			switch cn {
			case 0:
				barcode.Code = c.Value
			case 1:
				quantity, err := strconv.Atoi(c.Value)
				if err != nil {
					msg := " Row " + strconv.Itoa(rn+1) + "Quantity error"
					return errors.New(err.Error() + msg)
				}
				if quantity < 1 {
					msg := " Row " + strconv.Itoa(rn+1) + "Quantity error"
					return errors.New(err.Error() + msg)
				}
				barcode.Quantity = quantity
			case 2:
				barcode.SKU = c.Value
			}
			return err
		})
		barcodes = append(barcodes, barcode)
		return err
	})
	db := database.InitMySQL()
	repo := NewSettingRepository(db)
	err = repo.BatchCreateBarcode(barcodes)
	if err != nil {
		msg := "Batch creation error!"
		return errors.New(err.Error() + msg)
	}
	return nil
}

func (s *settingService) ExportBarcode(filter BarcodeFilterNoPage, path string) error {
	wb := xlsx.NewFile()
	sheet, err := wb.AddSheet("My New Sheet")
	if err != nil {
		msg := "Add Sheet Error!"
		return errors.New(err.Error() + msg)
	}
	headerRow := sheet.AddRow()
	header1 := headerRow.AddCell()
	header1.SetValue("Barcode")
	header2 := headerRow.AddCell()
	header2.SetValue("SKU")
	header3 := headerRow.AddCell()
	header3.SetValue("ItemName")
	header4 := headerRow.AddCell()
	header4.SetValue("Quantity")
	db := database.InitMySQL()
	repo := NewSettingRepository(db)
	list, err := repo.GetBarcodeListNoPage(filter)
	if err != nil {
		return err
	}
	for _, barcode := range list {
		row := sheet.AddRow()
		cell1 := row.AddCell()
		cell1.SetValue(barcode.Code)
		cell2 := row.AddCell()
		cell2.SetValue(barcode.SKU)
		cell3 := row.AddCell()
		cell3.SetValue(barcode.ItemName)
		cell4 := row.AddCell()
		cell4.SetValue(barcode.Quantity)
	}
	err = wb.Save(path)
	if err != nil {
		return err
	}
	sheet.Close()
	fmt.Println(path)
	return nil
}
