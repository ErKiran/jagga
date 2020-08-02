package utils

import (
	"math"
)

func HillBase(ropani float64, aana float64, paisa float64, daam float64, convertTo string) {

	totalAreaInSqMeter := ropani*ROPANI + aana*AANA + paisa*PAISA + daam*DAAM

	if convertTo == TERAISYSTEM {
		res := TeraiSystem(totalAreaInSqMeter)
		Pretty("res", res)
	}

	if convertTo == INTERNATIONALSYSTEM {
		res := InternationalSystem(totalAreaInSqMeter)
		Pretty("res", res)
	}

}

func TeraiBase(bigha float64, katha float64, dhur float64, convertTo string) {
	totalAreaInSqMeter := bigha*BIGHA + katha*KATHA + dhur*DHUR

	if convertTo == INTERNATIONALSYSTEM {
		res := InternationalSystem(totalAreaInSqMeter)
		Pretty("res", res)
	}

	if convertTo == HILLYSYSTEM {
		res := HillySystem(totalAreaInSqMeter)
		Pretty("res", res)
	}
}

func InternationalSystem(area float64) (result International) {
	return International{SQUAREMETER: area, SQUAREFOOT: area / SQUAREFOOT, HECTARE: area / HECTARE, ACRE: area / ACRE}
}

func HillySystem(area float64) (result Hill) {
	if area/ROPANI >= 1 {
		remainderAnna := math.Mod(area, ROPANI)
		if remainderAnna/AANA >= 1 {
			remainderPaisa := math.Mod(area, AANA)
			if remainderPaisa/PAISA >= 1 {
				remainderDaam := math.Mod(area, DAAM)
				res := Hill{ROPANI: math.Floor(area / ROPANI), AANA: math.Floor(remainderAnna / AANA), PAISA: math.Floor(remainderPaisa / PAISA), DAAM: remainderDaam / DAAM}
				return res
			}
			res := Hill{ROPANI: math.Floor(area / ROPANI), AANA: math.Floor(remainderAnna / AANA), DAAM: remainderPaisa / DAAM}
			return res
		}
		return
	}

	return
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
