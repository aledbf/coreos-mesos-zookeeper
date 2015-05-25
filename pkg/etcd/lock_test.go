package etcd

import (
	"testing"
	"time"
)

const (
	lockPath = "/lock-path"
)

func TestWaitForLock(t *testing.T) {
	startEtcd()
	defer stopEtcd()

	etcdClient := NewClient([]string{"http://localhost:4001"})

	err := AcquireLock(etcdClient, lockPath, 60)
	if err != nil {
		t.Fatalf("Unexpected err '%v'", err)
	}

	value := Get(etcdClient, lockPath)
	if value != "locked" {
		t.Fatalf("Expected '%v' arguments but returned '%v'", "locked", value)
	}

	err = ReleaseLock(etcdClient, lockPath)
	if err != nil {
		t.Fatalf("Unexpected error '%v'", err)
	}

	value = Get(etcdClient, lockPath)
	if value != "unlocked" {
		t.Fatalf("Expected '%v' arguments but returned '%v'", "unlocked", value)
	}

	err = WaitForLock(etcdClient, lockPath, 60, 10*time.Second)
	if err != nil {
		t.Fatalf("Unexpected error '%v'", err)
	}

	time.Sleep(10 * time.Second)

	value = Get(etcdClient, lockPath)
	if value != "locked" {
		t.Fatalf("Expected '%v' arguments but returned '%v'", "locked", value)
	}

	err = ReleaseLock(etcdClient, lockPath)
	if err != nil {
		t.Fatalf("Unexpected error '%v'", err)
	}
}
