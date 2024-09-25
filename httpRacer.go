package httpRacer

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(url1, url2 string, timeoutDuration time.Duration) (winner string, err error) {
	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(timeoutDuration):
		return "", fmt.Errorf("timeout occured after %v", timeoutDuration)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		_, err := http.Get(url)
		if err == nil {
			close(ch)
		}

	}()
	return ch
}
