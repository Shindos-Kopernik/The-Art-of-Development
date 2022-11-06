package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestExampleCleanup(t *testing.T) {
	Convey("Something works properly", t, func() {
		So(1, ShouldEqual, 1)
		So(2*2, ShouldEqual, 4)
		Convey("More test", func() { // t -  не передаем!!!
			So(1, ShouldEqual, 1)
			So(2*2, ShouldEqual, 4)
		})
	})
}
