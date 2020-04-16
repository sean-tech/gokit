package gokit

import (
	"fmt"
	"github.com/sean-tech/gokit/foundation"
	"github.com/sean-tech/gokit/logging"
	"sync"
	"testing"
)

func TestLog1(t *testing.T) {
	logging.Setup(logging.LogConfig{
		RunMode:         foundation.RUN_MODE_DEBUG,
		RuntimeRootPath: "/Users/lyra/",
		LogSavePath:     "Desktop/",
		LogPrefix:       "test",
	})
	for i := 0; i < 100000; i++ {
		logging.Debug(i)
	}
	fmt.Println("over")
}

var wg sync.WaitGroup
func TestGoroute(t *testing.T) {
	wg.Add(1)
	go func() {
		fmt.Println("done")
		wg.Done()
	}()
	print()

	wg.Add(1)
	go func() {
		fmt.Println("done")
		wg.Done()
	}()
	print()
}

func print()  {
	wg.Wait()
	fmt.Println("p")
}