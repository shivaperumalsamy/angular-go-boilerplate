package handlers

import (
	"fmt"
	"net/http"
)

// Hello is a simple Hello handler method
func Hello() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "<h1>Hello, World!</h1>\n")
	})
}
