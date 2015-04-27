package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestJsonEnv(t *testing.T) {
  os.Setenv("FOO", "fooof")
  os.Setenv("BAR", "snacks")

	Convey("Formatting a map of values as json", t, func() {
		data := map[string]string{
			"FOO": "foo-foo",
			"BAR": "snacks",
		}

		So(Json(data), ShouldEqual, "{\"BAR\":\"snacks\",\"FOO\":\"foo-foo\"}")
	})

	Convey("Formatting a map of values as Js", t, func() {
		data := map[string]string{
			"FOO": "bar",
		}
		So(Js(data), ShouldEqual, "window.ENV = {\"FOO\":\"bar\"};")
	})

	Convey("Extracting Values from the Enviroment", t, func() {
                envSize := len(os.Environ())

		Convey("Including Vars", func() {
			vars := []string{"FOO"}

			So(Env(vars, nil)["FOO"], ShouldEqual, "fooof")
			So(Env(vars, nil)["BAR"], ShouldEqual, "")
			So(len(Env(vars, nil)), ShouldEqual, 1)
		})

		Convey("Excluding Vars", func() {
			exclude := []string{"FOO"}

			So(Env(nil, exclude)["FOO"], ShouldEqual, "")
			So(Env(nil, exclude)["BAR"], ShouldEqual, "snacks")
                        So(len(Env(nil, exclude)), ShouldEqual, envSize - len(exclude))
		})

                Convey("Neither", func() {
			So(Env(nil, nil)["FOO"], ShouldEqual, "fooof")
			So(Env(nil, nil)["BAR"], ShouldEqual, "snacks")
                        So(len(Env(nil, nil)), ShouldEqual, envSize)
                })
	})

        Convey("Cli", t, func() {
          So(Cli([]string{"json_env", "--js", "-v", "FOO"}), ShouldEqual, "window.ENV = {\"FOO\":\"fooof\"};")
          So(Cli([]string{"json_env", "--json", "-v", "FOO"}), ShouldEqual, "{\"FOO\":\"fooof\"}")
        })
}
