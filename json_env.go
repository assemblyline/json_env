package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(Env())
}

func Env() string {
	env := make(map[string]string)

	for _, e := range strings.Split(os.Getenv("JSON_ENV"), ",") {
		env[e] = os.Getenv(e)
	}

	envJson, _ := json.Marshal(env)
	return string(envJson)
}
