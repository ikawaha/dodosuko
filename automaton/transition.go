package automaton

type input struct {
	status status
	input  string
}

var transition = map[input]status{
	{status: s0, input: "ドド"}: s1, {status: s0, input: "スコ"}: s0,
	{status: s1, input: "スコ"}: s2,
	{status: s2, input: "スコ"}: s3,
	{status: s3, input: "スコ"}: s4,

	{status: s4, input: "ドド"}: s5,
	{status: s5, input: "スコ"}: s6,
	{status: s6, input: "スコ"}: s7,
	{status: s7, input: "スコ"}: s8,

	{status: s8, input: "ドド"}:  s9,
	{status: s9, input: "スコ"}:  s10,
	{status: s10, input: "スコ"}: s11,
	{status: s11, input: "スコ"}: s12,
}
