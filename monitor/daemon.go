package monitor

import (
	"fmt"
	"time"
)

func Monitor() {
	fmt.Println("Monitoring...")
	ticker := time.NewTicker(time.Minute)
	go func() {
		for {
			for t := range ticker.C {
				fmt.Println("scan performed at ", t)
			}
		}
	}()
}
