package automaton

import (
	"context"
	"fmt"
	"io"
)

var debug bool //= true

func acceptor(ctx context.Context, ch <-chan string, w io.Writer) bool {
	q := start
	for {
		select {
		case v, ok := <-ch:
			if !ok {
				return false
			}
			q = transition[pair{state: q, input: v}]
			if debug && (q == start || q == s1) {
				fmt.Fprintln(w)
			}
			fmt.Fprint(w, v)
			if q == accept {
				fmt.Fprintln(w, "ラブ注入♡")
				return true
			}
		case <-ctx.Done():
			return false
		}
	}
}
