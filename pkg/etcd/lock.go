package etcd

import (
	"errors"
	"time"
)

func AcquireLock(c *Client, key string, ttl uint64) error {
	_, err := c.client.CompareAndSwap(key, "lock", ttl, "unlock", 0)
	return err
}

func ReleaseLock(c *Client, key string) error {
	_, err := c.client.Set(key, "unlock", 0)
	return err
}

func WaitForLock(c *Client, key string, ttl uint64, timeout time.Duration) error {
	successChan := make(chan bool)

	for {
		select {
		case <-successChan:
			return nil
		case <-getTimeoutChan(timeout):
			return errors.New("Timeout waiting for etcd lock")
		default:
			err := AcquireLock(c, key, ttl)
			if err == nil {
				successChan <- true
			}
		}
	}
}

// Return a channel that will send back a value when
// timeout.  If timeout is 0, the value will never arrive.
func getTimeoutChan(timeout time.Duration) <-chan time.Time {
	timeoutChan := make(<-chan time.Time)
	if timeout != 0 {
		timeoutChan = time.After(timeout)
	}
	return timeoutChan
}
