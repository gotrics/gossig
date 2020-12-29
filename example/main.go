// main.go
package main

import (
	"fmt"
	"os"
	"syscall"
	"time"

	"github.com/gotrics/gossig"
)

func HandlerSIGUSR1() {
	fmt.Println("SIGUSR1 received")
}

func main() {
	fmt.Printf("Test from console with:\nkill -SIGUSR1 %v\n", os.Getpid())
	gossig.SignalProcessor.HandlerAdd(syscall.SIGUSR1, HandlerSIGUSR1)
	fmt.Println("Start processing")
	gossig.SignalProcessor.Run()
	time.Sleep(time.Second * 20)
	fmt.Println("Stop processing")
	gossig.SignalProcessor.Stop()
	time.Sleep(time.Second * 5)
	fmt.Println("Restart processing")
	gossig.SignalProcessor.Run()
	time.Sleep(time.Second * 5)
}
