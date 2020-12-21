package dist

// Fn denotes a generic function to calculate a distance
type Fn func(reference, toMatch string) float64

// Generator denotes a generic distance metric generator
type Generator interface {

	// GenerateFn generates the actual matching distance function
	GenerateFn() Fn
}
