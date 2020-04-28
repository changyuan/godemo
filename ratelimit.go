package main

import "time"

const (
	rateLimit = time.Second / 100
)

type Client interface {
	Call(*Payload)
}

type Payload struct {
}

// func RateLimitCall(client Client, payloads []*Payload) {
// 	throttle := time.Tick(rateLimit)

// 	for _, payload := range payloads {
// 		<-throttle // rate limit our client calls
// 		go client.Call(payload)
// 	}
// }

// func BurstRateLimitCall(ctx context.Context, client Client, payloads []*Payload, burstLimit int) {
// 	throttle := make(chan time.Time, burstLimit)

// 	ctx, cancel := context.WithCancel(ctx)
// 	defer cancel()

// 	go func() {
// 		ticker := time.NewTicker(rateLimit)
// 		defer ticker.Stop()
// 		for t := range ticker.C {
// 			select {
// 			case throttle <- t:
// 			case <-ctx.Done():
// 				return // exit goroutine when surrounding function returns
// 			}
// 		}
// 	}()

// 	for _, payload := range payloads {
// 		<-throttle // rate limit our client calls
// 		go client.Call(payload)
// 	}
// }
