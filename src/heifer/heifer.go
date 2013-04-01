package main

import (
	"encoding/json"
	"os"
	"os/exec"
)

func yslow(url string) string {
	pwd, _ := os.Getwd()

	app := "phantomjs"
	arg0 := pwd + "/src/heifer/yslow.js"
	arg1 := "-icomps"
	arg2 := url

	cmd := exec.Command(app, arg0, arg1, arg2)
	out, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return ""
	}

	// Map of interfaces can receive any value types.
	u := map[string]interface{}{}
	json_err := json.Unmarshal(out, &u)
	if json_err != nil {
		println(json_err)
	}

	return string(out)
}

func main() {
	args := os.Args
	if len(args) > 1 {
		print(yslow(args[1]))
	} else {
		println("usage: " + os.Args[0] + " [url]")
	}
}
