package storage

// Storage is an interface to manage storage.
type Storage interface {
	Close() error

	GetItems() ([]Gif, error)

	GetItem(id int) (Gif, error)
	CreateItem(item Gif) (Gif, error)
	UpdateItem(id int, item Gif) (Gif, error)
	DeleteItem(id int) error
}
