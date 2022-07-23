package handlers

import (
	"fmt"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, Home")
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, About")
}
