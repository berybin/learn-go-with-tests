package selecty

import (
	"fmt"
	"net/http"
	"time"
)

// func Racer(a, b string) string {
// 	startA := time.Now()
// 	http.Get(a)
// 	aDuration := time.Since(startA)
//
// 	startB := time.Now()
// 	http.Get(b)
// 	bDuration := time.Since(startB)
//
// 	if aDuration < bDuration {
// 		return a
// 	}
//
// 	return b
// }

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out wating for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		defer close(ch)
		http.Get(url)
	}()

	return ch
}
