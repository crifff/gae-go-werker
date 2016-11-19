package module1

import (
	"gae-go-werker/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.Index1)
	http.Handle("/", r)
}
