package storage

// FakeStorage is Storage interface compatible structure for
// testing purposes.
type FakeStorage struct {
	Gifs      map[int]Gif
	CurrentID int
}

// Close method closes conection. For FakeStorage does nothing.
func (s *FakeStorage) Close() error {
	return nil
}

// GetItems returns all gifs in the storage.
func (s *FakeStorage) GetItems() ([]Gif, error) {
	result := make([]Gif, len(s.Gifs))

	i := 0
	for _, item := range s.Gifs {
		result[i] = item
		i++
	}

	return result, nil
}

// GetItem returns particular gif from the storage.
func (s *FakeStorage) GetItem(id int) (Gif, error) {
	result, ok := s.Gifs[id]

	if ok {
		return result, nil
	}
	return Gif{}, NotFoundError(id)
}

// CreateItem adds new gif to the storage if it passes validation.
func (s *FakeStorage) CreateItem(item Gif) (Gif, error) {
	item.ID = s.CurrentID

	if err := item.Validate(); err != nil {
		return Gif{}, err
	}

	s.Gifs[item.ID] = item

	s.CurrentID++

	return item, nil
}

// UpdateItem updates existing gif in the storage if it passes validation.
func (s *FakeStorage) UpdateItem(id int, item Gif) (Gif, error) {
	record, ok := s.Gifs[id]

	if !ok {
		return Gif{}, NotFoundError(id)
	}

	record.Update(item)

	if err := record.Validate(); err != nil {
		return Gif{}, err
	}

	s.Gifs[id] = record

	return record, nil
}

// DeleteItem removes specified item from the storage.
func (s *FakeStorage) DeleteItem(id int) error {
	if _, ok := s.Gifs[id]; !ok {
		return NotFoundError(id)
	}
	delete(s.Gifs, id)

	return nil
}

// NewFakeStorage creates a new storage object with populated gifs.
// For testing purposes.
func NewFakeStorage() (Storage, error) {
	return &FakeStorage{
		Gifs: map[int]Gif{
			1: Gif{ID: 1, Name: "Gif1", Slug: "gif-1", Width: 200, Height: 200},
			2: Gif{ID: 2, Name: "Gif2", Slug: "gif-2", Width: 200, Height: 200},
			3: Gif{ID: 3, Name: "Gif3", Slug: "gif-3", Width: 200, Height: 200},
			4: Gif{ID: 4, Name: "Gif4", Slug: "gif-4", Width: 200, Height: 200},
		},
		CurrentID: 5}, nil
}
