package httpRacer

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHttpRacer(t *testing.T) {
	t.Run("testing race - no timeout", func(t *testing.T) {
		slowServer := getConfigurableServer(10 * time.Millisecond)
		fastServer := getConfigurableServer(0 * time.Millisecond)
		want := fastServer.URL
		got, err := Racer(slowServer.URL, fastServer.URL, time.Millisecond*20)

		if err != nil {
			fmt.Print(err)
			t.Errorf("didn't expect error, got one")
		}
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("testing race - timeout", func(t *testing.T) {
		server1 := getConfigurableServer(5 * time.Millisecond)
		server2 := getConfigurableServer(6 * time.Millisecond)
		_, err := Racer(server1.URL, server2.URL, time.Millisecond*2)
		if err == nil {
			fmt.Print(err)
			t.Errorf("expected error, didn't get one")
		}
	})
}

func getConfigurableServer(waitTime time.Duration) (server *httptest.Server) {
	server = httptest.NewServer(http.HandlerFunc(getConfigurableCallBack(waitTime)))
	return server

}

func getConfigurableCallBack(waitTime time.Duration) func(w http.ResponseWriter, r *http.Request) {
	returnedFunc := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(waitTime)
		w.WriteHeader(http.StatusOK)
	}
	return returnedFunc
}
