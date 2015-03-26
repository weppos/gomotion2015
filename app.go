package main

import (
	"fmt"
	"os/exec"
	"strings"
	"net/http"
	"io/ioutil"

)

var (
	serverPort = "5000"
)


func main() {
	http.HandleFunc("/", RootHandler)

	fmt.Println(fmt.Sprintf("Listening on %s...", serverPort))
	err := http.ListenAndServe(":" + serverPort, nil)
	if err != nil {
		panic(err)
	}
}

func RootHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		fmt.Fprintln(res, "Usage: POST / <name>")
	case "POST":
		arg, _ := ioutil.ReadAll(req.Body)
		out, err := Dig(string(arg))
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
		} else {
			fmt.Fprintln(res, out)
		}
	default:
		http.NotFound(res, req)
	}
}

func Dig(arg string) (string, error) {
	fmt.Printf("Executing $ dig %s\n", arg)
	args := strings.Fields(arg)
	out, err := exec.Command("dig", args...).CombinedOutput()
	return string(out), err
}
