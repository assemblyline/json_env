package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

var format, vars string

func init() {
	flag.StringVar(&format, "f", "", "output format - js or json")
	flag.StringVar(&vars, "v", "", "The Environment Variables to expose")
	flag.Parse()
}

func main() {
	fmt.Println(Env(vars, format))
}

func Env(vars string, format string) string {
	switch strings.ToLower(format) {
	case "json":
		return jsonEnv(vars)
	case "js":
		return jsEnv(vars)
	}
	return ""
}

func jsonEnv(vars string) string {
	env := make(map[string]string)

	for _, e := range strings.Split(vars, ",") {
		env[e] = os.Getenv(e)
	}

	envJson, _ := json.Marshal(env)
	return string(envJson)
}

func jsEnv(vars string) string {
	return "window.ENV = " + jsonEnv(vars) + ";"
}
