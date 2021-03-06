package utils

type Terai struct {
	DHUR  float64 `json:"dhur"`
	KATHA float64 `json:"katha"`
	BIGHA float64 `json:"bigha"`
}

type Hill struct {
	DAAM   float64 `json:"daam"`
	PAISA  float64 `json:"paisa"`
	AANA   float64 `json:"aana"`
	ROPANI float64 `json:"ropani"`
}

type International struct {
	SQUAREMETER float64 `json:"m^2"`
	SQUAREFOOT  float64 `json:"ft^2"`
	HECTARE     float64 `json:"hectare"`
	ACRE        float64 `json:"acre"`
}

const (
	DAAM   = 1.99
	PAISA  = 7.95
	AANA   = 31.80
	ROPANI = 508.74

	DHUR  = 16.93
	KATHA = 338.62
	BIGHA = 6772.41

	HECTARE    = 10000
	ACRE       = 4046.86
	SQUAREFOOT = 0.092903

	HILLYSYSTEM         = "HILLYSYSTEM"
	TERAISYSTEM         = "TERAISYSTEM"
	INTERNATIONALSYSTEM = "INTERNATIONALSYSTEM"
)
