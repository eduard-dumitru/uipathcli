// Package utils contains command functionality which is reused across
// multiple other packages.
package utils

import "time"

const MaxAttempts = 3

// Retries the given function up to 3 times when it returns an RetryableError.
func Retry(f func() error) error {
	var err error
	for i := 1; ; i++ {
		err = f()
		_, retryable := err.(*RetryableError)
		if !retryable || i == MaxAttempts {
			return err
		}
		time.Sleep(1 * time.Second)
	}
}
