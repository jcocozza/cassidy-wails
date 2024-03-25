package utils

import (
	"encoding/json"
	"math"
)
// A custom stuct to allow for proper json marshalling to Inf and NaN float64's
type JsonFloat float64
func (value JsonFloat) MarshalJSON() ([]byte, error) {
	if math.IsNaN(float64(value)) || math.IsInf(float64(value), 0) {
		// If the value is NaN or Inf, return null
		return []byte("null"), nil
	}
	// Otherwise, use the default JSON encoding for float64
	return json.Marshal(float64(value))
}