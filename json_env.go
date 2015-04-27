package main

import (
	"encoding/json"
	"fmt"
	"github.com/jawher/mow.cli"
	"os"
	"strings"
)

func main() {
        fmt.Println(Cli(os.Args))
}

func Cli(args []string) string {
  jsonEnv := cli.App("json_env", "Expose ENV as json")
  jsonEnv.Spec = "(--js | --json) (-e...| -v...)"
  var (
    json    = jsonEnv.BoolOpt("json", false, "json format")
    js      = jsonEnv.BoolOpt("js", false, "js format")
    vars    = jsonEnv.StringsOpt("v vars", nil, "Envronment Variables to include")
    exclude = jsonEnv.StringsOpt("e exclude", nil, "Envronment Variables to exclude")
    output string
  )
  
  jsonEnv.Action = func() {
    env := Env(*vars, *exclude)
    if *json {
      output = Json(env)
    }
    if *js {
      output = Js(env)
    }
  }
  jsonEnv.Run(args)
  return output
}

func Json(env map[string]string) string {
	envJson, _ := json.Marshal(env)
	return string(envJson)
}

func Js(env map[string]string) string {
	return "window.ENV = " + Json(env) + ";"
}

func Env(vars []string, exclude []string) map[string]string {
	if len(vars) != 0 {
		return selectFromEnv(vars)
	}
	if len(exclude) != 0 {
		return envWithout(exclude)
	}
        return envWithout(nil)
}

func selectFromEnv(vars []string) map[string]string {
	env := make(map[string]string)
	for _, e := range vars {
		env[e] = os.Getenv(e)
	}
	return env
}

func envWithout(exclude []string) map[string]string {
	env := make(map[string]string)
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if !contains(exclude, pair[0]) {
			env[pair[0]] = pair[1]
		}
	}
	return env
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
