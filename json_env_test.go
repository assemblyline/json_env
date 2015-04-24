package main

import (
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func jsonEnv() map[string]string {
	var env map[string]string
	if err := json.Unmarshal([]byte(Env()), &env); err != nil {
		panic(err)
	}
	return env
}

func TestJsonEnv(t *testing.T) {
	Convey("json.Env", t, func() {

		Convey("When JSON_ENV is not set", func() {
			Convey("the JSON output is empty", func() {
				So(Env(), ShouldEqual, "{\"\":\"\"}")
			})

		})

		Convey("When JSON_ENV is set", func() {

			os.Setenv("JSON_ENV", "FOO,BAR")
			os.Setenv("FOO", "foo-foo")
			os.Setenv("BAR", "bar_snack")

			Convey("the JSON output contains the vars set", func() {
				So(jsonEnv()["FOO"], ShouldEqual, "foo-foo")
				So(jsonEnv()["BAR"], ShouldEqual, "bar_snack")
			})

			Convey("the JSON output includes unset vars as empty string values", func() {
				os.Setenv("JSON_ENV", "BAZ,boz")
				os.Setenv("boz", "boop")

				So(Env(), ShouldContainSubstring, "BAZ")
				So(jsonEnv()["BAZ"], ShouldEqual, "")
				So(Env(), ShouldContainSubstring, "boz\":\"boop")
				So(jsonEnv()["boz"], ShouldEqual, "boop")
			})

			Convey("does not include any vars not in JSON_ENV", func() {
				os.Setenv("BOOO", "NOT_HERE")
				So(Env(), ShouldNotContainSubstring, "BOOO")
				So(Env(), ShouldNotContainSubstring, "NOT_HERE")
			})

		})

	})
}
