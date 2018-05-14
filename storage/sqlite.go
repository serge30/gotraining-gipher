package storage

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type SqliteStorage struct {
	Db *gorm.DB
}

func (s *SqliteStorage) Close() error {
	return s.Db.Close()
}

func (s *SqliteStorage) GetItems() ([]Gif, error) {
	result := make([]Gif, 0)

	err := s.Db.Find(&result).Error

	return result, err
}

func (s *SqliteStorage) GetItem(id int) (Gif, error) {
	var item Gif
	if s.Db.First(&item, id).RecordNotFound() {
		return Gif{}, NotFoundError(id)
	}

	return item, nil
}

func (s *SqliteStorage) CreateItem(item Gif) (Gif, error) {
	if err := item.Validate(); err != nil {
		return Gif{}, err
	}

	if err := s.Db.Create(&item).Error; err != nil {
		return Gif{}, err
	}

	return item, nil
}

func (s *SqliteStorage) UpdateItem(id int, item Gif) (Gif, error) {
	var record Gif

	if err := s.Db.First(&record, id).Error; err != nil {
		return Gif{}, err
	}

	if item.Name != "" {
		record.Name = item.Name
	}

	if item.Slug != "" {
		record.Slug = item.Slug
	}

	if item.Width != 0 {
		record.Width = item.Width
	}

	if item.Height != 0 {
		record.Height = item.Height
	}

	if err := record.Validate(); err != nil {
		return Gif{}, err
	}

	if err := s.Db.Update(&record).Error; err != nil {
		return Gif{}, err
	}

	return record, nil
}

func (s *SqliteStorage) DeleteItem(id int) error {
	item := Gif{ID: id}

	return s.Db.Delete(&item).Error
}

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
