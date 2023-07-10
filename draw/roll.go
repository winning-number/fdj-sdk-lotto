package draw

// Roll is the roll(s) of a draw
// The second roll has been introduced for the V4.
type Roll struct {
	First     []int32
	Second    []int32
	LuckyBall int32
	HasLucky  bool
	HasSecond bool
}
