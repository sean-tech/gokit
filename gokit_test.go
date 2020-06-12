package gokit

import (
	"context"
	"fmt"
	"github.com/sean-tech/gokit/logging"
	"github.com/sean-tech/gokit/requisition"
	"sync"
	"testing"
)

var wg sync.WaitGroup

func TestLog1(t *testing.T) {
	logging.Setup(logging.LogConfig{
		LogSavePath:     "/Users/lyra/Desktop/",
		LogPrefix:       "test",
	})
	for i := 0; i < 100000; i++ {
		logging.Debug(i)
	}

	//wg.Add(100000)
	//for i := 0; i < 100000; i++ {
	//	go func(num int) {
	//		logging.Debug(num)
	//		wg.Done()
	//	}(i)
	//}
	//wg.Wait()
	fmt.Println("over")
}

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

func TestRequestion(t *testing.T) {
	ctx := requisition.NewRequestionContext(context.Background())
	requisition.GetRequisition(ctx).RequestId = 123456
	fmt.Print(requisition.GetRequisition(ctx))
}