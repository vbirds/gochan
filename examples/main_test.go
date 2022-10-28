package main

import (
	"fmt"
	"github.com/chalvern/gochan"
	"go.uber.org/zap"
	"testing"
	"time"
)

func TestMyMain(t *testing.T) {
	main()
}

func TestSingleGoroutine(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	sugar := logger.Sugar()
	gochan.SetLogger(sugar)

	gochanNum := 1
	bufferNum := 10
	manager := Manager{
		gochanNum:  gochanNum,
		bufferNum:  bufferNum,
		dispatcher: gochan.NewDispatcher(gochanNum, bufferNum),
	}

	//objID := 1
	var myNumber int = 0
	for i := 0; i < 100000; i++ {
		task1 := func() error {
			myNumber++
			return nil
		}
		manager.Dispatch(-1, task1)
	}

	time.Sleep(time.Second * 2)
	if myNumber != 100000 {
		fmt.Printf("myNumber:%d\n", myNumber)
		panic("myNumber should be 1000")
	}
}
