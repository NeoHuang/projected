package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/NeoHuang/projected/handlers"
)

const (
	usageMessage = `A Webhook for github to assign newly opened issue to a configured project`
)

func init() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, usageMessage)

		flag.PrintDefaults()
	}
}

func main() {
	port, _, _, _, _ := parseFlags()

	http.Handle("/", &handlers.ProjectedHandler{})

	log.Printf("Starting server on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Panicf("Can't start server (%s)", err)
	}
}

func parseFlags() (port, defaultProject, secretToken, username, token string) {
	flag.StringVar(&port, "port", "80", "http service port")
	flag.StringVar(&defaultProject, "project", "", "which project should newly opened issue assigned to")
	flag.StringVar(&secretToken, "secret", "", "secret token for github webhook")
	flag.StringVar(&username, "username", "", "github user name")
	flag.StringVar(&token, "token", "", "github personal access token")
	flag.Parse()

	return port, defaultProject, secretToken, username, token
}
