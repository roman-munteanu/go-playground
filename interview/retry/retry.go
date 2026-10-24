package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func main() {
	fmt.Println("retry")
}

func Retry(
	ctx context.Context,
	attempts int,
	initialDelay time.Duration,
	maxDelay time.Duration,
	operation func(c context.Context) error) error {

	if attempts < 1 {
		return errors.New("enter a valid number of attempts")
	}

	if initialDelay < 0 || maxDelay < 0 {
		return errors.New("delays cannot be negative")
	}

	if maxDelay < initialDelay {
		maxDelay = initialDelay
	}

	delay := initialDelay

	for attempt := 0; attempt < attempts; attempt++ {

		if err := ctx.Err(); err != nil {
			return err
		}

		err := operation(ctx)
		if err == nil {
			return nil
		}

		if attempt == attempts-1 {
			return err
		}

		timer := time.NewTimer(delay)

		select {
		case <-ctx.Done():
			if !timer.Stop() {
				select {
				case <-timer.C:
				default:
				}
			}
			return ctx.Err()
		case <-timer.C:
		}

		if delay >= maxDelay/2 {
			delay = maxDelay
		} else {
			delay *= 2
		}

	}

	return nil
}
