package wait

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func WaitForServer(url string) error {
	timeout := 1 * time.Minute
	deadlines := time.Now().Add(timeout)

	for tries := 0; time.Now().Before(deadlines) ; tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}

		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries))
	}

	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
