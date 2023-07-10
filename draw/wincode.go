package draw

// WinCode is the win code(s) of a draw
// Each ticket has a unique code. From V3, at the end of the draw, a code(s) are drawn.
type WinCode struct {
	Codes  []string
	Price  float64
	Number int
}
