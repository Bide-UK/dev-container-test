package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/alive", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "true")
	})

	http.ListenAndServe(":8000", nil)
}
