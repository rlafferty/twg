package main

import (
	"net/http"

	"github.com/rlafferty/twg/app"
)

func main() {
	http.ListenAndServe(":3000", &app.Server{})
}
