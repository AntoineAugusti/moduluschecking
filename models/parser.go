package models

// A parser should be able to:
// - read the weights, exception information and algorithm to use for sort code ranges.
// - give a list of sort code substitutions.
type Parser interface {
	// Read the weights, exception information and algorithm to use for sort code ranges.
	Weights() map[string]SortCodeData
	// Give a list of sort code substitutions
	Substitutions() map[string]string
}
