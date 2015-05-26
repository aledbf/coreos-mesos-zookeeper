package net

import (
	"testing"
	"time"
)

func TestListenTCP(t *testing.T) {
	err := WaitForPort("tcp", "127.0.0.1", RandomPort("tcp"), 5*time.Second)
	if err != nil {
		t.Fatal(err)
	}
}

func TestListenUDP(t *testing.T) {
	err := WaitForPort("udp", "127.0.0.1", RandomPort("udp"), 5*time.Second)
	if err != nil {
		t.Fatal(err)
	}
}
