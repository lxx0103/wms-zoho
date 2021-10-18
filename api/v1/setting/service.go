package setting

import (
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
	//Barcode Management
	GetBarcodeByID(int64) (Barcode, error)
	NewBarcode(BarcodeNew) (Barcode, error)
	GetBarcodeList(BarcodeFilter) (int, []Barcode, error)
	UpdateBarcode(int64, BarcodeNew) (Barcode, error)
	GetBarcodeByCode(string) (*Barcode, error)
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
