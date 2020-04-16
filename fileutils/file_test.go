package fileutils

import (
	"fmt"
	"testing"
	"time"
)

func TestCopyFile(t *testing.T) {

	var src = "/Users/lyra/Desktop/f1.txt"
	var dst = "/Users/lyra/Desktop/f1_copied.txt"
	var written int64
	var err error
	if written, err = CopyFile(dst, src); err != nil {
		t.Error(err)
		return
	}
	fmt.Println("written is : ", written)

	time.Sleep(3*time.Second)
	if err = ClearFile(src); err != nil {
		t.Error(err)
		return
	}
}