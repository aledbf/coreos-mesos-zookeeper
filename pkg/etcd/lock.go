package etcd

import (
	etcdlock "github.com/leeor/etcd-sync"
)

func AcquireLock(c *Client, key string, ttl uint64) error {
	c.lock = etcdlock.NewMutexFromClient(c.client, key, ttl)
	return c.lock.Lock()
}

func ReleaseLock(c *Client, key string) {
	c.lock.Unlock()
}
