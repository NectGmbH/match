package normalize

// Fn denotes a generic string normalization function
type Fn func(string) string

// Fns denotes a list of normalization functions
type Fns []Fn
