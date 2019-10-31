package service

import (
	"fmt"
	"net/http"
)

type echo struct{}

func (e *echo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "echo\n")
}
