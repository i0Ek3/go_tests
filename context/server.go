package main

import (
	"context"
	"fmt"
	_ "log"
	"net/http"
)

// Store fetchs the data
type Store interface {
	Fetch(ctx context.Context) (string, error)
}

// Server returns a handler
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		data, err := store.Fetch(ctx)

		if err != nil {
			//log.Print("== met err ----> ", err)
			return
		}

		fmt.Fprint(w, data)
	}
}
