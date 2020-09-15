package app

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}
func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)
	return mux
}
