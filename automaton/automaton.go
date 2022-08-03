package automaton

import (
	"context"
	"os"
)

func Run(ctx context.Context) bool {
	ch := make(chan string, 1)
	go randomInput(ctx, ch)
	return acceptor(ctx, ch, os.Stdout)
}
