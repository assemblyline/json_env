package main

import (
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func parseEnv(vars string) map[string]string {
	var env map[string]string
	if err := json.Unmarshal([]byte(Env(vars, "Json")), &env); err != nil {
		panic(err)
	}
	return env
}

func TestJsonEnv(t *testing.T) {
	Convey("Env", t, func() {

		Convey("an unsupported format", func() {
			So(Env("FOO", "sausages"), ShouldEqual, "")
		})

		Convey("The js format", func() {
			os.Setenv("FOO", "foo-foo")

			Convey("the output wraps the json in some javascript", func() {
				So(Env("FOO", "js"), ShouldEqual, "window.ENV = {\"FOO\":\"foo-foo\"};")
				So(Env("FOO", "JS"), ShouldEqual, "window.ENV = {\"FOO\":\"foo-foo\"};")
			})
		})
		Convey("the json format", func() {

			Convey("When vars is not set", func() {
				Convey("the JSON output is empty", func() {
					So(jsonEnv(""), ShouldEqual, "{\"\":\"\"}")
				})

			})

			Convey("When the vars is set", func() {

				os.Setenv("FOO", "foo-foo")
				os.Setenv("BAR", "bar_snack")

				Convey("the JSON output contains the vars set", func() {
					So(parseEnv("FOO,BAR")["FOO"], ShouldEqual, "foo-foo")
					So(parseEnv("FOO,BAR")["BAR"], ShouldEqual, "bar_snack")
				})

				Convey("the JSON output includes unset vars as empty string values", func() {
					os.Setenv("boz", "boop")

					So(Env("BAZ,boz", "json"), ShouldContainSubstring, "BAZ")
					So(parseEnv("BAZ,boz")["BAZ"], ShouldEqual, "")
					So(Env("BAZ,boz", "json"), ShouldContainSubstring, "boz\":\"boop")
					So(parseEnv("BAZ,boz")["boz"], ShouldEqual, "boop")
				})

				Convey("does not include any vars not in vars", func() {
					os.Setenv("BOOO", "NOT_HERE")
					So(Env("", "json"), ShouldNotContainSubstring, "BOOO")
					So(Env("", "json"), ShouldNotContainSubstring, "NOT_HERE")
				})

			})

		})
	})

}
