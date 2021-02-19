package tests

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"testing"
)

func TestIndex(t *testing.T) {

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			res, err := http.Get("http://localhost:8090")
			if res == nil || err != nil {
				t.Error("Err")
			}
			buf := new(strings.Builder)
			io.Copy(buf, res.Body)
			fmt.Println(buf)
		}()
	}

	wg.Wait()
}
