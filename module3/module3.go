package module3

import (
	"gae-go-werker/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.Index3)
	http.Handle("/", r)
}
