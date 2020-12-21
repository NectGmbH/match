package dist

import "github.com/neuneck/smetrics"

const (
	defaultWeightInsert  = 1
	defaultWeightReplace = 1
	defaultWeightAdd     = 1
)

// WagnerFischer denotes a fuzzy matcher based on the Wagner-Fischer algorithm
type WagnerFischer struct {
	weightInsert  int
	weightReplace int
	weightAdd     int
}

// NewWagnerFischer instantiates a new WagnerFischer matcher
func NewWagnerFischer() *WagnerFischer {
	return &WagnerFischer{
		weightInsert:  defaultWeightInsert,
		weightReplace: defaultWeightReplace,
		weightAdd:     defaultWeightAdd,
	}
}

// Weights sets the weights / penalties for the Levenshtein distance calculation
func (w *WagnerFischer) Weights(weightInsert, weightReplace, weightAdd int) *WagnerFischer {
	w.weightInsert = weightInsert
	w.weightReplace = weightReplace
	w.weightAdd = weightAdd

	return w
}

// GenerateFn generates the WagnerFischer matching distance function
func (w *WagnerFischer) GenerateFn() Fn {
	return func(reference, toMatch string) float64 {
		return float64(smetrics.WagnerFischer(reference, toMatch, w.weightAdd, w.weightInsert, w.weightReplace))
	}
}
