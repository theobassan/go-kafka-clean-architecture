package goroutinemanager

import (
	"context"
	"fmt"
	"syscall"
	"time"

	"github.com/x-mod/routine"
)

func forLoop(ctx context.Context) error {
	i := 0
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			fmt.Println("forLoop", i)
			i += 1
			time.Sleep(2 * time.Second)
		}
	}
}

func Run() {
	ctx, cancel := context.WithCancel(context.Background())
	routine.Main(
		ctx,
		//main Go
		routine.ExecutorFunc(forLoop),
		//signals
		routine.Signal(syscall.SIGINT, routine.SigHandler(func() {
			cancel()
			//os.Exit(1)
		})),
	)
}
