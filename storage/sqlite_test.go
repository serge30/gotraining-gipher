package storage

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSqliteStorage(t *testing.T) {
	Convey("Given SqliteStorage", t, func() {
		storage, err := NewSqliteStorage(":memory:")
		storage.(*SqliteStorage).testData()
		defer storage.Close()

		Convey("No error on storage creation", func() {
			So(err, ShouldBeNil)
		})

		Convey("When listed", func() {
			gifs, err := storage.GetItems()

			Convey("Should be no error", func() {
				So(err, ShouldBeNil)
			})

			Convey("Should be 4 items", func() {
				So(gifs, ShouldHaveLength, 4)
			})
		})

		Convey("When get existing item", func() {
			gif, err := storage.GetItem(2)

			Convey("Should be no error", func() {
				So(err, ShouldBeNil)
			})

			Convey("Should have correct ID", func() {
				So(gif.ID, ShouldEqual, 2)
			})
		})

		Convey("When get missing item", func() {
			gif, err := storage.GetItem(5)

			Convey("Should be NotFound error", func() {
				So(err, ShouldBeError, "Id 5 is not found")
			})

			Convey("Should return zero Gif", func() {
				So(gif, ShouldBeZeroValue)
			})
		})

		Convey("When create valid gif", func() {
			gif := Gif{Name: "Gif5", Slug: "gif-5", Width: 300, Height: 200}
			resultGif, err := storage.CreateItem(gif)

			Convey("Should be no error", func() {
				So(err, ShouldBeNil)
			})

			Convey("Should return Gif with ID 5", func() {
				So(resultGif.ID, ShouldEqual, 5)
			})

			Convey("Should item with ID 5 present in storage", func() {
				gif, err := storage.GetItem(5)

				So(err, ShouldBeNil)
				So(gif.ID, ShouldEqual, 5)
			})
		})

		Convey("When create invalid gif", func() {
			gif := Gif{Name: "", Slug: "gif-5", Width: 300, Height: 200}
			resultGif, err := storage.CreateItem(gif)

			Convey("Should be ValidationError", func() {
				So(err, ShouldBeError, "Field 'name' is malformed or missing")
			})

			Convey("Should return zero Gif", func() {
				So(resultGif.ID, ShouldBeZeroValue)
			})

			Convey("Should be item with ID 5 absent in storage", func() {
				_, err := storage.GetItem(5)

				So(err, ShouldBeError, "Id 5 is not found")
			})
		})

		Convey("When update by valid gif", func() {
			gif := Gif{Name: "Gif3-new", Slug: "gif-3-new", Width: 1300, Height: 1200}
			resultGif, err := storage.UpdateItem(3, gif)

			Convey("Should be no error", func() {
				So(err, ShouldBeNil)
			})

			Convey("Should return Gif with ID 3", func() {
				So(resultGif.ID, ShouldEqual, 3)
			})

			Convey("Should item with ID 3 present in storage with updated fields", func() {
				gif, err := storage.GetItem(3)

				So(err, ShouldBeNil)
				So(gif, ShouldResemble, Gif{ID: 3, Name: "Gif3-new", Slug: "gif-3-new", Width: 1300, Height: 1200})
			})
		})

		Convey("When update missing item", func() {
			gif := Gif{Name: "Gif3-new", Slug: "gif-3-new", Width: 1300, Height: 1200}
			resultGif, err := storage.UpdateItem(5, gif)

			Convey("Should be NotFoundError error", func() {
				So(err, ShouldBeError, "Id 5 is not found")
			})

			Convey("Should return zero Gif", func() {
				So(resultGif, ShouldBeZeroValue)
			})

			Convey("Should item with ID 5 be still missing", func() {
				_, err := storage.GetItem(5)

				So(err, ShouldBeError, "Id 5 is not found")
			})
		})

		Convey("When update by invalid gif", func() {
			gif := Gif{Name: "Gif3-new", Slug: "gif-3-new", Width: -1300, Height: 1200}
			resultGif, err := storage.UpdateItem(3, gif)

			Convey("Should be ValidationError error", func() {
				So(err, ShouldBeError, "Field 'width' is malformed or missing")
			})

			Convey("Should return zero Gif", func() {
				So(resultGif, ShouldBeZeroValue)
			})

			Convey("Should item with ID 3 be unmodified", func() {
				gif, err := storage.GetItem(3)

				So(err, ShouldBeNil)
				So(gif, ShouldResemble, Gif{ID: 3, Name: "Gif3", Slug: "gif-3", Width: 200, Height: 200})
			})
		})

		Convey("When delete existing item", func() {
			err := storage.DeleteItem(3)

			Convey("Should be no error", func() {
				So(err, ShouldBeNil)
			})

			Convey("Should item with ID 3 be removed", func() {
				_, err := storage.GetItem(3)
				So(err, ShouldBeError, "Id 3 is not found")
			})
		})

		Convey("When delete missing item", func() {
			err := storage.DeleteItem(5)

			Convey("Should be NotFoundError error", func() {
				So(err, ShouldBeError, "Id 5 is not found")
			})
		})
	})
}

func (s *SqliteStorage) testData() error {
	gifs := []Gif{
		Gif{ID: 1, Name: "Gif1", Slug: "gif-1", Width: 200, Height: 200},
		Gif{ID: 2, Name: "Gif2", Slug: "gif-2", Width: 200, Height: 200},
		Gif{ID: 3, Name: "Gif3", Slug: "gif-3", Width: 200, Height: 200},
		Gif{ID: 4, Name: "Gif4", Slug: "gif-4", Width: 200, Height: 200},
	}

	for _, gif := range gifs {
		if err := s.Db.Create(&gif).Error; err != nil {
			return err
		}

	}

	return nil
}
