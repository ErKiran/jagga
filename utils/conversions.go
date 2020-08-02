package utils

import (
	"math"
)

func HillBase(ropani float64, aana float64, paisa float64, daam float64, convertTo string) {

	totalInSqMeter := ropani*ROPANI + aana*AANA + paisa*PAISA + daam*DAAM

	if convertTo == TERAISYSTEM {
		res := TeraiSystem(totalInSqMeter)
		Pretty("res", res)
	}

	if convertTo == INTERNATIONALSYSTEM {
		res := InternationalSystem(totalInSqMeter)
		Pretty("res", res)
	}

}

func InternationalSystem(area float64) (result International) {
	return International{SQUAREMETER: area, SQUAREFOOT: area / SQUAREFOOT, HECTARE: area / HECTARE, ACRE: area / ACRE}
}

func TeraiSystem(area float64) (result Terai) {
	if area/BIGHA >= 1 {
		remainderKatha := math.Mod(area, BIGHA)
		if remainderKatha/KATHA >= 1 {
			remainderDhur := math.Mod(area, KATHA)
			res := Terai{BIGHA: math.Floor(area / BIGHA), KATHA: math.Floor(remainderKatha / KATHA), DHUR: remainderDhur / DHUR}
			return res
		}
		res := Terai{BIGHA: math.Floor(area / BIGHA), KATHA: 0, DHUR: remainderKatha / DHUR}
		return res
	}

	if area/KATHA >= 1 {
		unit := math.Mod(area, KATHA)
		res := Terai{BIGHA: 0, KATHA: math.Floor(area / KATHA), DHUR: unit / DHUR}
		return res
	}

	res := Terai{DHUR: area / DHUR}
	return res
}
