package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func TimeoutMiddleware(timeout time.Duration, next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx, cancel := context.WithTimeout(r.Context(), timeout)
		defer cancel() 

		r = r.WithContext(ctx)

		done := make(chan struct{})
		go func() {
			t := time.Now()
			next.ServeHTTP(w, r)
			fmt.Printf("Time taken by Request- %v\n", time.Since(t))
			close(done)
			
		}()

		select {
		case <-done:	
		case <-ctx.Done():
			if ctx.Err() == context.DeadlineExceeded {
				http.Error(w, "Request timed out", http.StatusGatewayTimeout)
			}
		}
	})
}
