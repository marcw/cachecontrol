package cachecontrol

import (
	"testing"
	"time"
)

func TestParsing(t *testing.T) {
	cc := Parse("public, max-age=10")

	if !cc.Public() {
		t.Error("Public() should return true")
	}
	if 10*time.Second != cc.MaxAge() {
		t.Error("MaxAge() should return a time.Duration of 10 seconds")
	}
}
