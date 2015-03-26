package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

var (
	serverPort = "5000"
)

func main() {
	http.HandleFunc("/", RootHandler)

	log.Println(fmt.Sprintf("Listening on %s...", serverPort))

	err := http.ListenAndServe(":"+serverPort, nil)
	if err != nil {
		panic(err)
	}
}

func RootHandler(res http.ResponseWriter, req *http.Request) {
	log.Printf("Started %s %s\n", req.Method, req.RequestURI)
	defer log.Printf("Completed %s %s\n", req.Method, req.RequestURI)

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
	log.Printf("Executing $ dig %s\n", arg)
	args := strings.Fields(arg)
	out, err := exec.Command("dig", args...).CombinedOutput()
	return string(out), err
}
