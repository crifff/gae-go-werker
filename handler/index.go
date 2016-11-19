package handler

import (
	"fmt"
	"net/http"
)

// Index1 is root page for module1
func Index1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "this page is module1")
}

// Index2 is root page for module2
func Index2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "this page is module2")
}

// Index3 is root page for module3
func Index3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "this page is module3")
}
