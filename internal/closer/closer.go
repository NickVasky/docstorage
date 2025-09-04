package closer

import (
	"log"
	"os"
	"os/signal"
	"sync"
)

var globalCloser = New()

// Public *interface*
func Add(f ...func() error) {
	globalCloser.add(f...)
}

func CloseAll() {
	globalCloser.closeAll()
}

func Wait() {
	globalCloser.wait()
}

type Closer struct {
	mu    sync.Mutex
	once  sync.Once
	done  chan struct{}
	funcs []func() error
}

func New(sig ...os.Signal) *Closer {
	c := &Closer{done: make(chan struct{})}

	if len(sig) > 0 {
		go func() {
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, sig...)
			<-ch
			signal.Stop(ch)
			c.closeAll()
		}()
	}
	return c
}

func (c *Closer) wait() {
	<-c.done
}

func (c *Closer) add(f ...func() error) {
	c.mu.Lock()
	c.funcs = append(c.funcs, f...)
	c.mu.Unlock()
}

func (c *Closer) closeAll() {
	c.once.Do(
		func() {
			defer close(c.done)

			c.mu.Lock()
			funcs := c.funcs
			c.funcs = nil
			c.mu.Unlock()

			errs := make(chan error, len(funcs))
			for _, f := range funcs {
				go func(f func() error) {
					errs <- f()
				}(f)
			}

			for i := 0; i < cap(errs); i++ {
				if err := <-errs; err != nil {
					log.Printf("Closer error: %v", err)
				}
			}
		},
	)
}
