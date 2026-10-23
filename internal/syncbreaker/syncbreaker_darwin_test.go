package syncbreaker

import (
	"testing"
	"time"
)

func TestBreak(t *testing.T) {
	// First ensure clean state
	Break()

	// Create semaphore
	if !CreateTestSemaphore() {
		t.Fatal("failed to create test semaphore")
	}

	// Break should successfully destroy it
	if !Break() {
		t.Error("expected Break() to succeed destroying the semaphore")
	}

	// Second break should fail since it's already gone
	if Break() {
		t.Error("expected second Break() to return false")
	}
}

func TestWatcher(t *testing.T) {
	Break()

	watcher := StartWatcher(50 * time.Millisecond)
	defer watcher.Stop()

	// Create semaphore while watcher is running
	if !CreateTestSemaphore() {
		t.Fatal("failed to create test semaphore")
	}

	// Wait for watcher to clear it
	time.Sleep(150 * time.Millisecond)

	if Break() {
		t.Error("expected semaphore to have been unlinked by watcher")
	}
}
