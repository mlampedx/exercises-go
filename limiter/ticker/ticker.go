package ticker

import "time"

func Create(limiter chan<- time.Time) {
	for t := range time.Tick(200 * time.Millisecond) {
		limiter <- t
	}
}
