package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/luhn", getLuhn)

	err := http.ListenAndServe(":3333", mux)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("error starting serverL %s\n", err)
		os.Exit(1)
	}
}
