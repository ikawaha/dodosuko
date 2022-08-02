package main

import (
	"context"
	"os"
	"time"

	"github.com/ikawaha/dodosuko/automaton"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	if !automaton.Run(ctx) {
		os.Exit(1)
	}
}
