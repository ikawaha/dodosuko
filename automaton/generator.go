package automaton

import (
	"context"
	"math/rand"
	"time"
)

func gen(ctx context.Context, ch chan<- string) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	v := []string{"ドド", "スコ"}
	for {
		select {
		case <-ctx.Done():
			return
		default:
			ch <- v[r.Intn(len(v))]
		}
	}
}
