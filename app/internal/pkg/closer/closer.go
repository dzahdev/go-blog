package closer

import (
	"log"
	"os"
	"os/signal"
	"sync"
)

var globalCloser = New()

func Append(f ...func() error) {
	globalCloser.Append(f...)
}

func Wait() {
	globalCloser.Wait()
}

func CloseAll() {
	globalCloser.CloseAll()
}

type Closer struct {
	mu    sync.Mutex
	once  sync.Once
	done  chan struct{}
	funcs []func() error
}

// New creates a new Closer. If sig is provided, it listens for the signal and closes all the registered functions.
func New(sig ...os.Signal) *Closer {
	c := &Closer{done: make(chan struct{})}

	if len(sig) > 0 {
		go func() {
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, sig...)
			<-ch
			signal.Stop(ch)
			c.CloseAll()
		}()
	}

	return c
}

// Append appends a function to the closer. The function will be called when the closer is closed.
func (c *Closer) Append(f ...func() error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.funcs = append(c.funcs, f...)
}

// Wait waits for the closer to be closed. It blocks until the closer is closed.
func (c *Closer) Wait() {
	<-c.done
}

// CloseAll closes all the registered functions.
func (c *Closer) CloseAll() {
	c.once.Do(func() {
		defer close(c.done)
		c.mu.Lock()
		funcs := c.funcs
		c.funcs = nil
		c.mu.Unlock()

		// Run all the functions in separate goroutines.
		errs := make(chan error, len(funcs))
		for _, f := range funcs {
			go func(f func() error) {
				errs <- f()
			}(f)
		}

		for i := 0; i < len(errs); i++ {
			if err := <-errs; err != nil {
				log.Println("error returned from Closer")
			}
		}
	})
}
