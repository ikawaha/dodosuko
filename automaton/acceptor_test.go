package automaton

import (
	"bytes"
	"context"
	"testing"
)

func Test_acceptor(t *testing.T) {
	testdata := []struct {
		name   string
		input  []string
		want   bool
		output string
	}{
		{
			name:   "exact match",
			input:  []string{"ドド", "スコ", "スコ", "スコ", "ドド", "スコ", "スコ", "スコ", "ドド", "スコ", "スコ", "スコ"},
			want:   true,
			output: "ドドスコスコスコドドスコスコスコドドスコスコスコラブ注入♡\n",
		},
		{
			name: "postfix match",
			input: []string{
				"スコ", "スコ", "ドド", "スコ", "スコ", "ドド",
				"ドド", "スコ", "スコ", "スコ", "ドド", "スコ", "スコ", "スコ", "ドド", "スコ", "スコ", "スコ"},
			want:   true,
			output: "スコスコドドスコスコドド" + "ドドスコスコスコドドスコスコスコドドスコスコスコラブ注入♡\n",
		},
		{
			name: "prefix match",
			input: []string{
				"ドド", "スコ", "スコ", "スコ", "ドド", "スコ", "スコ", "スコ", "ドド", "スコ", "スコ", "スコ",
				"スコ", "スコ", "ドド",
			},
			want:   true,
			output: "ドドスコスコスコドドスコスコスコドドスコスコスコラブ注入♡\n",
		},
		{
			name:   "not accept",
			input:  []string{"ドド", "スコ", "スコ", "スコ", "ドド", "スコ", "スコ", "スコ", "ドド", "スコ", "スコ", "ドド"},
			want:   false,
			output: "ドドスコスコスコドドスコスコスコドドスコスコドド",
		},
		{
			name: "not accept: s4 -> s0",
			input: []string{
				"ドド", "スコ", "スコ", "スコ", // s4 -> s0
				"スコ", "スコ", "スコ", "スコ", "ドド", "スコ", "スコ", "スコ", "ドド", "スコ", "スコ", "スコ",
			},
			want:   false,
			output: "ドドスコスコスコ" + "スコスコスコスコドドスコスコスコドドスコスコスコ",
		},
		{
			name: "not accept: s8 -> s0",
			input: []string{
				"ドド", "スコ", "スコ", "スコ", "ドド", "スコ", "スコ", "スコ", // s8 -> s0
				"スコ", "スコ", "スコ", "スコ", "ドド", "スコ", "スコ", "スコ", "ドド", "スコ", "スコ", "スコ",
			},
			want:   false,
			output: "ドドスコスコスコドドスコスコスコ" + "スコスコスコスコドドスコスコスコドドスコスコスコ",
		},
	}
	for _, tt := range testdata {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.TODO())
			defer cancel()
			ch := make(chan string, 1)
			go func(ctx context.Context, input []string) {
				select {
				default:
					for _, v := range input {
						ch <- v
					}
					close(ch)
				case <-ctx.Done():
					close(ch)
				}
			}(ctx, tt.input)
			var b bytes.Buffer
			if got := acceptor(context.TODO(), ch, &b); got != tt.want {
				t.Errorf("want %t, but got %t", tt.want, got)
			}
			if got, want := b.String(), tt.output; got != want {
				t.Errorf("want %s, got %s", want, got)
			}
		})
	}
}
