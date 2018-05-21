package storage

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGif(t *testing.T) {
	Convey("Given valid Gif struct", t, func() {
		gif := Gif{ID: 1, Name: "Gif1", Slug: "gif-1", Width: 300, Height: 200}

		Convey("When Validation method is run", func() {
			err := gif.Validate()

			Convey("No error should be", func() {
				So(err, ShouldBeNil)
			})
		})

		Convey("When String method is run", func() {
			str := gif.String()

			Convey("The value should be correct", func() {
				So(str, ShouldEqual, "GIF<1: Gif1>")
			})
		})
	})
	Convey("Given invalid Gif struct", t, func() {
		gif := Gif{ID: 1, Name: "Gif1", Slug: "gif-1", Width: 300, Height: 200}

		Convey("When Name is empty and Validation method is run", func() {
			gif.Name = ""
			err := gif.Validate()

			Convey("Should be validation error with Name field", func() {
				So(err, ShouldBeError, "Field 'name' is malformed or missing")
			})
		})

		Convey("When Slug is empty and Validation method is run", func() {
			gif.Slug = ""
			err := gif.Validate()

			Convey("Should be validation error with Slug field", func() {
				So(err, ShouldBeError, "Field 'slug' is malformed or missing")
			})
		})

		Convey("When Width is zero and Validation method is run", func() {
			gif.Width = 0
			err := gif.Validate()

			Convey("Should be validation error with Width field", func() {
				So(err, ShouldBeError, "Field 'width' is malformed or missing")
			})
		})

		Convey("When Height is zero and Validation method is run", func() {
			gif.Height = 0
			err := gif.Validate()

			Convey("Should be validation error with Height field", func() {
				So(err, ShouldBeError, "Field 'height' is malformed or missing")
			})
		})

	})
}
