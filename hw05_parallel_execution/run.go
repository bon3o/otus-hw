package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var (
	ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
	ErrNoWorkersPassed     = errors.New("no workers were assigned")
	ErrNegativeWorkerCount = errors.New("negative number of workers")
)

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if n == 0 {
		return ErrNoWorkersPassed
	}
	if n < 0 {
		return ErrNegativeWorkerCount
	}
	if m <= 0 {
		return ErrErrorsLimitExceeded
	}
	taskCh := make(chan Task)
	var wg sync.WaitGroup
	var errCnt int32

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for task := range taskCh {
				if err := task(); err != nil {
					atomic.AddInt32(&errCnt, 1)
				}
			}
		}()
	}
	for _, task := range tasks {
		if atomic.LoadInt32(&errCnt) >= int32(m) {
			break
		}
		taskCh <- task
	}
	close(taskCh)

	wg.Wait()

	if errCnt >= int32(m) {
		return ErrErrorsLimitExceeded
	}

	return nil
}
