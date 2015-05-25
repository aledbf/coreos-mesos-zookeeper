package etcd

import (
	"errors"
	"time"
)

func AcquireLock(c *Client, key string, ttl uint64) error {
	_, err := c.client.CompareAndSwap(key, "locked", ttl, "unlocked", 0)
	if err != nil {
		etcdErr := convertEtcdError(err)
		// if key not found, lock is free
		if etcdErr.ErrorCode == 100 {
			log.Debugf("creating key %s to hold the lock", key)
			_, err := c.client.Set(key, "locked", ttl)
			return err
		}
	}

	return err
}

func ReleaseLock(c *Client, key string) error {
	log.Debugf("releasing lock key %s", key)
	_, err := c.client.Set(key, "unlocked", 0)
	if err != nil {
		log.Debugf("error releasing lock key %s: %v", key, err)
	}

	return err
}

func WaitForLock(c *Client, key string, ttl uint64, timeout time.Duration) error {
	log.Debugf("waiting for the lock creation in %s", key)

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
			} else {
				log.Debugf("error creating lock in key %s '%v'", key, err)
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
