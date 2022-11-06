package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestExampleCleanup(t *testing.T) {
	x := 0
	Convey("A", t, func() {
		x++
		Convey("A-B", func() { // t -  не передаем!!!
			x++ // TODO Лучшее место для Setup. Здесь можно подготавливать объекты для тестирования
			Convey("A-B-C1", func() {
				So(x, ShouldEqual, 2)
			})
			Convey("A-B-C2", func() {
				So(x, ShouldEqual, 4)
			})
			Convey("A-B-C3", func() {
				panic("IM IN PANIC")
			})
		})
		Reset(func() {
			t.Log("Finish")
		})
	})
}
