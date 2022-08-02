package automaton

import (
	"context"
)

func Run(ctx context.Context) bool {
	ch := make(chan string, 1)
	go gen(ctx, ch)
	return acceptor(ctx, ch)
}
