package foundation

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestGenerateId(t *testing.T) {
	fmt.Println(time.Now().Unix())

	var testData = []struct{
		workerId int64
	} {
		{0},
		{1},
		{2},
		{3},
		{4},
	}

	for _, data := range testData {
		worker, err := NewWorker(data.workerId)
		if err != nil {
			t.Errorf(err.Error())
		}
		fmt.Println(worker.GetId())
	}
}

func TestGenerateId2(t *testing.T) {
	var ids []int64 = []int64{}
	var lock sync.Mutex
	wg := sync.WaitGroup{}

	// 注意，worker全局应只有一个实例，不要生成id时临时创建，并发会冲突
	worker, err := NewWorker(1)
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			lock.Lock()
			defer lock.Unlock()
			ids = append(ids, worker.GetId())
		}()
	}
	wg.Wait()
	fmt.Printf("%v", ids)
}
