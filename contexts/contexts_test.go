package contexts

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	t.Run("returns data from the store", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, t: t} // to take out T?
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("got '%s', want '%s'", response.Body.String(), data)
		}
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, t: t}
		svr := Server(store) // this just creates the handler function that performs a particular action e.g. in this case on the Store resource

		request := httptest.NewRequest(http.MethodGet, "/", nil) //simulates a get request

		cancellingCtx, cancel := context.WithCancel(request.Context()) //creating a new / duplicate context with a cancel function
		time.AfterFunc(5*time.Millisecond, cancel)                     // cancel that context after some time for this test
		request = request.WithContext(cancellingCtx)                   // assign request to the new and soon to be cancelled context

		response := &SpyResponseWriter{}

		svr.ServeHTTP(response, request) //calls the handler function using the response and request variables that suit the handler's function signature
		//The above function is what actually calls the handler function

		// //assert that activity was cancelled
		if response.written {
			t.Error("a response should not have been written")
		}

	})
}

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() { // separate goroutine to perform the action
		var result string
		for _, c := range s.response {
			select { //helps wait for a cancellation, otherwise, continue performing the action
			case <-ctx.Done():
				log.Println("spy store got cancelled") //while performing the action, if cancelled will notify and return from the for loop i.e. the activity
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select { //main goroutine waits for either the final result or a cancellation
	case <-ctx.Done(): // has to return something even if cancelled, so will return empty string and the reason for cancellation
		return "", ctx.Err() // this is what deals with return, while the first select deals with breaking out of the activity
	case res := <-data:
		return res, nil

	}

}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}
