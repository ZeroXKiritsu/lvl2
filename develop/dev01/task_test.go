package main

import (
	"testing"
	"time"
)

const host = "0.beevik-ntp.pool.ntp.org"

func TestAccurateTime(t *testing.T) {
	tm, _ := accurateTime(host)
	now := time.Now()
	t.Run("IsNil", func(t *testing.T) {
		t.Logf("Local Time %v\n", now)
		t.Logf("~True Time %v\n", tm)
		t.Logf("Offset %v\n", tm.Sub(now))
	})
}