package models

// Give information about a sort code
type SortCodeData struct {
	// The algorithm to use:
	// - DBLAL: double alternate check
	// - MOD10: modulus with a modulus of 10
	// - MOD11: modulus with a modulus of 11
	Algorithm string
	// Weights to use for the sort code and the account number
	Weights []int
	// The exception value. If 0, the sort code does not have an exception.
	ExceptionValue int
	// If a sort code follows multiple rule, Next will not be nil.
	Next *SortCodeData
	// Indicates the line number where this rule was stored in the file.
	LineNumber int
}

// HasException Does this sort code has got an exception?
func (s SortCodeData) HasException() bool {
	return s.ExceptionValue > 0
}

// HasNext Does this sort code follows another rule?
func (s SortCodeData) HasNext() bool {
	return s.Next != nil
}

// IsException checks if a sort code follows an exception.
func (s SortCodeData) IsException(exceptionValue int) bool {
	return s.HasException() && s.ExceptionValue == exceptionValue
}

// FollowsExceptions checks if a sort code follows 2 exceptions.
func (s SortCodeData) FollowsExceptions(ex1, ex2 int) bool {
	is1 := s.IsException(ex1)
	has2 := s.HasNext() && s.Next.IsException(ex2)

	return is1 && has2
}
