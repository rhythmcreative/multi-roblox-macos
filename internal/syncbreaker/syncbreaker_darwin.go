package syncbreaker

/*
#include <semaphore.h>
#include <fcntl.h>
#include <stdio.h>

int destroy_semaphore() {
    const char *sem_name = "/RobloxPlayerUniq";
    if (sem_unlink(sem_name) == -1)
    {
        return 1;
    }
    return 0;
}

int create_semaphore() {
    sem_t *sem = sem_open("/RobloxPlayerUniq", O_CREAT, 0644, 1);
    if (sem == SEM_FAILED) {
        return 1;
    }
    return 0;
}
*/
import "C"
import (
	"sync"
	"time"
)

// Break attempts to unlink the RobloxPlayerUniq semaphore.
// Returns true if the semaphore was found and unlinked, false otherwise.
func Break() bool {
	return C.destroy_semaphore() == 0
}

// CreateTestSemaphore is used by unit tests to create a semaphore.
func CreateTestSemaphore() bool {
	return C.create_semaphore() == 0
}

// Pulse repeatedly unlinks the semaphore every interval for the given duration.
// This ensures that during Roblox launch, the semaphore is cleared as soon as it's registered.
func Pulse(duration time.Duration, interval time.Duration) {
	go func() {
		end := time.Now().Add(duration)
		for time.Now().Before(end) {
			Break()
			time.Sleep(interval)
		}
	}()
}

// Watcher runs a background monitor that periodically unlinks the semaphore.
type Watcher struct {
	stopCh chan struct{}
	once   sync.Once
}

func StartWatcher(interval time.Duration) *Watcher {
	w := &Watcher{
		stopCh: make(chan struct{}),
	}
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-w.stopCh:
				return
			case <-ticker.C:
				Break()
			}
		}
	}()
	return w
}

func (w *Watcher) Stop() {
	w.once.Do(func() {
		close(w.stopCh)
	})
}
