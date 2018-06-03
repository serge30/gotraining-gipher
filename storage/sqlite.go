package storage

import (
	"github.com/jinzhu/gorm"

	// Support for Sqlite DB.
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// SqliteStorage is Storage interface compatible structure which
// provides storage using Sqlite as backend.
type SqliteStorage struct {
	Db *gorm.DB
}

// Close method closes conection to DB.
func (s *SqliteStorage) Close() error {
	return s.Db.Close()
}

// GetItems returns all gifs in the storage.
func (s *SqliteStorage) GetItems() ([]Gif, error) {
	result := make([]Gif, 0)

	err := s.Db.Find(&result).Error

	return result, err
}

// GetItem returns particular gif from the storage.
func (s *SqliteStorage) GetItem(id int) (Gif, error) {
	var item Gif
	if s.Db.First(&item, id).RecordNotFound() {
		return Gif{}, NotFoundError(id)
	}

	return item, nil
}

// CreateItem adds new gif to the storage if it passes validation.
func (s *SqliteStorage) CreateItem(item Gif) (Gif, error) {
	if err := item.Validate(); err != nil {
		return Gif{}, err
	}

	if err := s.Db.Create(&item).Error; err != nil {
		return Gif{}, err
	}

	return item, nil
}

// UpdateItem updates existing gif in the storage if it passes validation.
func (s *SqliteStorage) UpdateItem(id int, item Gif) (Gif, error) {
	var record Gif

	if s.Db.First(&record, id).RecordNotFound() {
		return Gif{}, NotFoundError(id)
	}

	record.Update(item)

	if err := record.Validate(); err != nil {
		return Gif{}, err
	}

	if err := s.Db.Save(&record).Error; err != nil {
		return Gif{}, err
	}

	return record, nil
}

// DeleteItem removes specified item from the storage.
func (s *SqliteStorage) DeleteItem(id int) error {
	item := Gif{ID: id}

	result := s.Db.Delete(&item)
	rowsAffected, err := result.RowsAffected, result.Error

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return NotFoundError(id)
	}

	return nil
}

// NewSqliteStorage connects to specifies Sqlite DB file and returns
// a new storage interface.
func NewSqliteStorage(dbFileName string) (Storage, error) {
	db, err := gorm.Open("sqlite3", dbFileName)
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&Gif{}).Error
	if err != nil {
		return nil, err
	}

	return &SqliteStorage{Db: db}, nil
}
