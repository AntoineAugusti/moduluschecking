package models

// This structure describes a sort code range, starting from a
// sort code and ending to another.
type SortCodeRange struct {
	// Beginning of the sort code range.
	Start int
	// End of the sort code range.
	End int
	// The algorithm to use:
	// - DBLAL: double alternate check
	// - MOD10: modulus with a modulus of 10
	// - MOD11: modulus with a modulus of 11
	Algorithm string
	// Weights to use for the sort code and the account number
	Weights []int
	// The exception value. If 0, the sort code does not have an exception.
	ExceptionValue int
	// Indicates the line number where this rule was stored in the file.
	LineNumber int
}

// HasException checks if a sort code range has got an exception
func (sc SortCodeRange) HasException() bool {
	return sc.ExceptionValue > 0
}
