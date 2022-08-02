package automaton

import (
	"context"
	"fmt"
)

func acceptor(ctx context.Context, ch <-chan string) bool {
	q := start
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				return false
			}
			q = transition[input{status: q, input: v}]
			fmt.Print(v)
			if q == finite {
				fmt.Println("ラブ注入♡")
				return true
			}
		case <-ctx.Done():
			return false
		}
	}
}
