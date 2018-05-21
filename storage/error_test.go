package storage

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestErrors(t *testing.T) {
	Convey("Given NotFoundError", t, func() {
		testError := NotFoundError(5)

		Convey("When error message is taken", func() {
			msg := testError.Error()

			Convey("The message should be as defined", func() {
				So(msg, ShouldEqual, "Id 5 is not found")
			})
		})
	})
	Convey("Given ValidationError", t, func() {
		testError := ValidationError("test")

		Convey("When error message is taken", func() {
			msg := testError.Error()

			Convey("The message should be as defined", func() {
				So(msg, ShouldEqual, "Field 'test' is malformed or missing")
			})
		})
	})
}
