package automaton

type pair struct {
	state state
	input string
}

var transition = map[pair]state{
	{state: s0, input: "ドド"}: s1, {state: s0, input: "スコ"}: s0,
	{state: s1, input: "スコ"}: s2,
	{state: s2, input: "スコ"}: s3,
	{state: s3, input: "スコ"}: s4,

	{state: s4, input: "ドド"}: s5, {state: s5, input: "スコ"}: s0,
	{state: s5, input: "スコ"}: s6,
	{state: s6, input: "スコ"}: s7,
	{state: s7, input: "スコ"}: s8,

	{state: s8, input: "ドド"}: s9, {state: s9, input: "スコ"}: s0,
	{state: s9, input: "スコ"}:  s10,
	{state: s10, input: "スコ"}: s11,
	{state: s11, input: "スコ"}: s12,
}
