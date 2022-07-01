package main

import "net"
import "net/http"
import "fmt"
import "os/exec"
import "runtime"

type MyHandler struct {
}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}


func launchInBrowser(url string) {
	var prog string
	var args []string
	if runtime.GOOS == "windows" {
		prog = "explorer"
		args = []string{url}
	} else if runtime.GOOS == "darwin" {
		prog = "open"
		args = []string{url}
	} else {
		prog = "x-www-browser"
		args = []string{url}
	}
	browserCmd := exec.Command(prog, args...)

	if err := browserCmd.Run(); err != nil {
		panic(err)
	}

}

func main() {
	port := "8080"
	url := "http://localhost:8080"

	l, err := net.Listen("tcp", ":" + port)

	if err != nil {
		// handle error
	}

	// launch browser
	launchInBrowser(url)

	handler := MyHandler{}
	if err := http.Serve(l, handler); err != nil {
		// handle error
	}

}
