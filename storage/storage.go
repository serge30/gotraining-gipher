package storage

type Storage interface {
	GetItems() ([]Gif, error)

	GetItem(id int) (Gif, error)
	CreateItem(item Gif) (Gif, error)
	UpdateItem(id int, item Gif) (Gif, error)
	DeleteItem(id int) error
}
