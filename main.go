package main

import "net"
import "net/http"
import "fmt"
import "os/exec"
import "runtime"
import "github.com/CCDirectLink/ccms2/internal/downloader"
func apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request", r.URL.String())
	downloader.FromGithubRepo("https://github.com/2hh8899/ArcaneLab/tree/dev")

	fmt.Fprintf(w, "Wassup");
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the base url.");
}


func createApiServer() *http.ServeMux {
	sm := http.NewServeMux()
	sm.HandleFunc("/", baseHandler);
	sm.HandleFunc("/api/", apiHandler)
	
	return sm
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

	serv := createApiServer()

	if err := http.Serve(l, serv); err != nil {
		// handle error
	}

}
