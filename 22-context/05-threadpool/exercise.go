package threadpool

import (
	"context"
	"fmt"
	"sync"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// Runnable represents a task that can be run with a context
type Runnable interface {
	Run(context.Context) error
}

// ThreadPool interface definition
type ThreadPool interface {
	Run(Runnable)
	Close()
}

// threadPool implementation
type threadPool struct {
	tasks   chan Runnable      // Pufferelt csatorna a feladatoknak
	errChan chan error         // Pufferelt csatorna a hibáknak
	ctx     context.Context    // A pool globális contextje (leállításhoz)
	cancel  context.CancelFunc // Leállító függvény
	wg      sync.WaitGroup     // Munkások számontartása
	once    sync.Once          // A Close idempotenciájához
}

// NewThreadPool creates a new threadpool with n workers
func NewThreadPool(n int) (ThreadPool, chan error) {
	ctx, cancel := context.WithCancel(context.Background())
	
	// A hibacsatornát puffereljük, hogy ne blokkolja a workereket
	errChan := make(chan error, 100)

	// FONTOS JAVÍTÁS: A feladatok csatornáját is puffereljük!
	// Így a Run() nem blokkol, ha a munkások épp dolgoznak.
	tasks := make(chan Runnable, 100)

	tp := &threadPool{
		tasks:   tasks,
		errChan: errChan,
		ctx:     ctx,
		cancel:  cancel,
	}

	// Elindítunk n darab munkást (worker)
	for i := 0; i < n; i++ {
		tp.wg.Add(1)
		go tp.worker()
	}

	return tp, errChan
}

// A munkás (worker) logikája
func (tp *threadPool) worker() {
	defer tp.wg.Done()
	for {
		select {
		// 1. Ha a poolt lezárták, kilépünk
		case <-tp.ctx.Done():
			return

		// 2. Várjuk a feladatokat
		case task, ok := <-tp.tasks:
			if !ok {
				return
			}
			// Futtatjuk a feladatot, átadva neki a pool contextjét
			if err := task.Run(tp.ctx); err != nil {
				// Hibakezelés: próbáljuk beküldeni, de nem blokkolunk
				select {
				case tp.errChan <- err:
				default:
					// Ha tele a csatorna, stdout-ra írunk (Requirement hint)
					fmt.Printf("ThreadPool error channel full, dropped error: %v\n", err)
				}
			}
		}
	}
}

// Run submits a task to the threadpool
func (tp *threadPool) Run(task Runnable) {
	select {
	case tp.tasks <- task:
		// Sikeresen beküldtük a feladatot a pufferbe
	case <-tp.ctx.Done():
		// Ha a pool már le van zárva, nem csinálunk semmit
		return
	}
}

// Close stops all running threads and closes the error channel
func (tp *threadPool) Close() {
	// A sync.Once biztosítja, hogy csak egyszer fusson le a leállítás
	tp.once.Do(func() {
		// 1. Minden munkásnak szólunk, hogy álljon le
		tp.cancel()

		// 2. Egy külön goroutine-ban várjuk meg a leállást
		go func() {
			tp.wg.Wait()      // Megvárjuk, míg minden worker kilép
			close(tp.errChan) // Csak utána zárjuk le a hibacsatornát
		}()
	})
}
