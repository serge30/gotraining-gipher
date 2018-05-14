package storage

type FakeStorage struct {
	Gifs      map[int]Gif
	CurrentID int
}

func (s *FakeStorage) Close() error {
	return nil
}

func (s *FakeStorage) GetItems() ([]Gif, error) {
	result := make([]Gif, len(s.Gifs))

	i := 0
	for _, item := range s.Gifs {
		result[i] = item
		i++
	}

	return result, nil
}

func (s *FakeStorage) GetItem(id int) (Gif, error) {
	result, ok := s.Gifs[id]

	if ok {
		return result, nil
	}
	return Gif{}, NotFoundError(id)
}

func (s *FakeStorage) CreateItem(item Gif) (Gif, error) {
	item.ID = s.CurrentID

	if err := item.Validate(); err != nil {
		return Gif{}, err
	}

	s.Gifs[item.ID] = item

	s.CurrentID++

	return item, nil
}

func (s *FakeStorage) UpdateItem(id int, item Gif) (Gif, error) {
	record, ok := s.Gifs[id]

	if !ok {
		return Gif{}, NotFoundError(id)
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

	s.Gifs[id] = record

	return record, nil
}

func (s *FakeStorage) DeleteItem(id int) error {
	if _, ok := s.Gifs[id]; !ok {
		return NotFoundError(id)
	}
	delete(s.Gifs, id)

	return nil
}

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
