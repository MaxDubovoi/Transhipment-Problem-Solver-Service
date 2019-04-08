package data


type ComputeDistributionRequest struct {
	Providers []float64   `json:"providers"`
	Consumers []float64   `json:"consumers"`
	Costs     [][]float64 `json:"costs"`
}