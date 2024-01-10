package circuit_breaker

import (
	"fmt"
	"time"

	"github.com/rosariocannavo/api_gateway/internal/nats"
	"github.com/sony/gobreaker"
)

var CircuitBreaker *gobreaker.CircuitBreaker

func init() {
	CircuitBreaker = gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        "Circuit Breaker",
		MaxRequests: 0,               // Maximum number of consecutive failures before tripping the circuit
		Interval:    5 * time.Second, // Duration to wait before allowing another request
		Timeout:     2 * time.Second, // Timeout for a single request attempt
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures > 3
		},
		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
			fmt.Printf("Circuit breaker '%s' changed from '%s' to '%s'\n", name, from, to)

			message := fmt.Sprintf("Timestamp: %s | Circuit breaker '%s' changed from '%s' to '%s'\n", time.Now().UTC().Format(time.RFC3339), name, from, to)
			nats.NatsConnection.PublishMessage(message)
		},
	})
}
