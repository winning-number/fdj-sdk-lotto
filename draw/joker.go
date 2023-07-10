package draw

// Joker is the joker number of lotto
// The plus number has already existed in the past.
// The base number has introduced for the V1 then removed for the V2.
type Joker struct {
	Base    string
	Plus    string
	HasBase bool
}
