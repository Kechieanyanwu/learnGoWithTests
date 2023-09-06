package contexts

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context()) //the actual function is to return the data inside the resource
		if err != nil {
			log.Println("context was cancelled")
			return
		}
		fmt.Fprint(w, data) //then write this through the response writer

	}
}

type Store interface {
	Fetch(ctx context.Context) (string, error) //all activities are sub-tasks that are performed within a particular context, so helps with cancellations
}

// our interface takes in a context, and returns a string and an optional error in case the context was cancelled
