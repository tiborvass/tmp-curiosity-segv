package main

// #include <unistd.h>
import "C"

import (
	"context"

	"golang.org/x/sys/unix"
)

// contextWithCCancel will return a context.Context, but where the CancelFunc
// is a C.int (really a FD for a POSIX pipe(2)), where when written to, will
// cancel the function.
func contextWithCCancel(ctx context.Context) (context.Context, C.int, error) {
	ctx, cancel := context.WithCancel(ctx)

	pipes := [2]int{}
	if err := unix.Pipe(pipes[:]); err != nil {
		return nil, -1, err
	}

	go func() {
		defer cancel()

		b := make([]byte, 1)
		unix.Read(pipes[0], b)
		unix.Close(pipes[0])
		unix.Close(pipes[1])
	}()

	return ctx, C.int(pipes[1]), nil
}
