package maths

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	)

func TestApp(t *testing.T){
	Convey("it should sum 2 numbers", t, func() {
		So(Sum(2,3), ShouldEqual,5)
	})

}